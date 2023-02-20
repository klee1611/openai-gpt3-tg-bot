package openaichatbot_test

import (
	"os"
	"testing"

	chat "github.com/klee1611/openai-tg-bot"
)

func TestOpenAI(t *testing.T) {
	envSetup()
	defer envTearDown()

	t.Run("TestOpenAINormalResponse", func(t *testing.T) {
		mockOpenAIServer := mockOpenAIServer()
		defer mockOpenAIServer.Close()

		openAIAPI := chat.OpenAIAPI{
			BaseURL: mockOpenAIServer.URL,
			Token:   os.Getenv("OPENAI_API_KEY"),
		}

		res, err := openAIAPI.OpenAI("Test request message")
		if err != nil {
			t.Errorf("openAI() error = %v", err)
		}
		if res != "Hello, world!" {
			t.Errorf("openAI() = %v, want %v", res, "Hello, world!")
		}
	})
}
