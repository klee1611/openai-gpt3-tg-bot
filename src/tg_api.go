package openaigpt3bot

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type TelegramAPI struct {
	BaseURL string
	Token   string
}

func (t *TelegramAPI) SendTgMsg(
	chatID float64,
	message string,
) error {
	sendMsgURL := fmt.Sprintf("%s/bot%s/sendMessage", t.BaseURL, t.Token)
	reqBody := map[string]interface{}{
		"chat_id": chatID,
		"text":    message,
	}

	log.Info("Sending messages to telegram...")
	resp, err := client.R().
		SetBody(reqBody).
		SetHeader("Content-Type", "application/json").
		SetResult(map[string]interface{}{}).
		Post(sendMsgURL)
	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Errorf("Error response from telegram send message: %v", resp)
		return errors.New("error response from telegram send message")
	}
	log.Info("Message sent to telegram")
	return nil
}
