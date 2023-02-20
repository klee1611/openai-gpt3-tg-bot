package openaichatbot_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	chat "github.com/klee1611/openai-tg-bot"
)

func wrapper(body string, tgToken bool) (int, error) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(body))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	if tgToken {
		req.Header.Set(
			"X-Telegram-Bot-Api-Secret-Token",
			os.Getenv("TG_HEADER_TOKEN"),
		)
	}

	rr := httptest.NewRecorder()
	chat.TgWebHook(rr, req)

	return rr.Code, nil
}

func TestTgWebHook(t *testing.T) {
	envSetup()
	defer envTearDown()

	mockOpenAIServer := mockOpenAIServer()
	defer mockOpenAIServer.Close()
	os.Setenv("OPENAI_API_URL", mockOpenAIServer.URL)

	mockTgServer := mockTGServer()
	defer mockTgServer.Close()
	os.Setenv("TG_API_URL", mockTgServer.URL)

	t.Run("TestTgWebHookSuccess", func(t *testing.T) {
		body := `{
			"message": {
				"text": "Test request message",
				"message_id": 1234567890,
				"chat": {
					"id": 1234567890
				},
				"from": {
					"id": 1234567890
				}
			}
		}`
		code, err := wrapper(body, true)
		if err != nil {
			t.Errorf("TgWebHook() error = %v", err)
		}
		if code != http.StatusOK {
			t.Errorf("TgWebHook() = %v, want %v", code, http.StatusOK)
		}
	})

	t.Run("TestTgWebHookNoToken", func(t *testing.T) {
		body := `{
			"message": {
				"text": "Test request message",
				"message_id": 1234567890,
				"chat": {
					"id": 1234567890
				},
				"from": {
					"id": 1234567890
				}
			}
		}`
		code, err := wrapper(body, false)
		if err != nil {
			t.Errorf("TgWebHook() error = %v", err)
		}
		if code != http.StatusUnauthorized {
			t.Errorf("TgWebHook() = %v, want %v", code, http.StatusUnauthorized)
		}
	})

	t.Run("TestTgWebHookWrongUser", func(t *testing.T) {
		body := `{
			"message": {
				"text": "Test request message",
				"message_id": 1234567890,
				"chat": {
					"id": 1234567890
				},
				"from": {
					"id": 1234567891
				}
			}
		}`
		code, err := wrapper(body, true)
		if err != nil {
			t.Errorf("TgWebHook() error = %v", err)
		}
		if code != http.StatusUnauthorized {
			t.Errorf("TgWebHook() = %v, want %v", code, http.StatusUnauthorized)
		}
	})
}
