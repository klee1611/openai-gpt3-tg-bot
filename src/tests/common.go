package openaigpt3bot_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
)

// TG Response
type From struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

type tgResponse struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	ErrorCode   int     `json:"error_code"`
	Description string  `json:"description"`
}

// OpenAI Response
type openAIResponse struct {
	ID      string
	Object  string
	Created int
	Model   string
	Choices []struct {
		Text string
	}
}

// Response
var (
	tgNormalResponse = tgResponse{
		Ok: true,
		Result: Message{
			MessageID: 1234567890,
			From: From{
				ID:        1234567890,
				IsBot:     false,
				FirstName: "Test",
				Username:  "test",
			},
			Chat: Chat{
				ID:        1234567890,
				FirstName: "Test",
				LastName:  "Test",
				Username:  "test",
				Type:      "private",
			},
			Date: 1234567890,
			Text: "Test message",
		},
	}
	openAINormalResponse = openAIResponse{
		ID:      "cmpl-1234567890",
		Object:  "text_completion",
		Created: 1234567890,
		Model:   "text-davinci-003",
		Choices: []struct {
			Text string
		}{
			{
				Text: "Hello, world!",
			},
		},
	}
)

// Environment variables for testing
func envSetup() {
	os.Setenv("TG_BOT_TOKEN", "test_tg_bot_token")
	os.Setenv("TG_API_URL", "test_tg_api_url")
	os.Setenv("OPENAI_API_URL", "test_openai_api_url")
	os.Setenv("OPENAI_API_KEY", "test_openai_api_key")
	os.Setenv("TG_HEADER_TOKEN", "test_tg_header_token")
	os.Setenv("TG_USER_ID", "1234567890")
}

func envTearDown() {
	os.Unsetenv("TG_BOT_TOKEN")
	os.Unsetenv("TG_API_URL")
	os.Unsetenv("OPENAI_API_URL")
	os.Unsetenv("OPENAI_API_KEY")
	os.Unsetenv("TG_HEADER_TOKEN")
	os.Unsetenv("TG_USER_ID")
}

// Mock server for testing
func mockOpenAIServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(openAINormalResponse)
			},
		),
	)
}

func mockTGServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(tgNormalResponse)
			},
		),
	)
}
