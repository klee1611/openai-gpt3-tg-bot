package openaigpt3bot

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var (
	client *resty.Client = resty.New()
)

func init() {
	functions.HTTP("TgWebHook", TgWebHook)
}

type HookRequest struct {
	Message struct {
		Chat struct {
			ID float64 `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
		From struct {
			ID float64 `json:"id"`
		} `json:"from"`
	} `json:"message"`
}

func TgWebHook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	tgHeaderToken, ok := r.Header["X-Telegram-Bot-Api-Secret-Token"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if len(tgHeaderToken) != 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if tgHeaderToken[0] != os.Getenv("TG_HEADER_TOKEN") {
		log.Errorf("Invalid telegram header token: %v", tgHeaderToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Info("Valid tg header")

	reqBody := HookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("Invalid request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: Add decode error response
		return
	}
	reqUserID := reqBody.Message.From.ID

	if tgUserID, err := strconv.ParseFloat(os.Getenv("TG_USER_ID"), 64); err != nil {
		log.Errorf("Invalid telegram user id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if reqUserID != tgUserID {
		log.Errorf("Invalid telegram user id: %v", reqUserID)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Info("Valid telegram user")

	gpt3API := OpenAIGPT3API{
		BaseURL: os.Getenv("OPENAI_API_URL"),
		Token:   os.Getenv("OPENAI_API_KEY"),
	}
	resp, err := gpt3API.GPT3(reqBody.Message.Text)
	if err != nil {
		log.Printf("Error: OpenAI GPT3 API error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: Add openAI GPT3 API error response
		return
	}

	tg := TelegramAPI{
		BaseURL: os.Getenv("TG_API_URL"),
		Token:   os.Getenv("TG_BOT_TOKEN"),
	}
	if err := tg.SendTgMsg(reqBody.Message.Chat.ID, resp); err != nil {
		log.Printf("Error: telegram send message error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: Add telegram send message error response
		return
	}
	w.WriteHeader(http.StatusOK)
}
