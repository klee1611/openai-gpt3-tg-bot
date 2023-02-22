package openaigpt3bot_test

import (
	"os"
	"testing"

	gpt3 "github.com/klee1611/openai-gpt3-tg-bot"
)

func TestSendTgMsg(t *testing.T) {
	envSetup()
	defer envTearDown()

	t.Run("TestSendTgMsg", func(t *testing.T) {
		mockTgServer := mockTGServer()
		defer mockTgServer.Close()

		tgAPI := gpt3.TelegramAPI{
			BaseURL: mockTgServer.URL,
			Token:   os.Getenv("TG_BOT_TOKEN"),
		}

		err := tgAPI.SendTgMsg(1234567890, "Test request message")
		if err != nil {
			t.Errorf("sendTgMsg() error = %v", err)
		}
	})
}
