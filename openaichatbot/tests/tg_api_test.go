package openaichatbot_test

import (
	"os"
	"testing"

	chat "github.com/klee1611/openai-tg-bot"
)

func TestSendTgMsg(t *testing.T) {
	envSetup()
	defer envTearDown()

	t.Run("TestSendTgMsg", func(t *testing.T) {
		mockTgServer := mockTGServer()
		defer mockTgServer.Close()

		tgAPI := chat.TelegramAPI{
			BaseURL: mockTgServer.URL,
			Token:   os.Getenv("TG_BOT_TOKEN"),
		}

		err := tgAPI.SendTgMsg(1234567890, "Test request message")
		if err != nil {
			t.Errorf("sendTgMsg() error = %v", err)
		}
	})
}
