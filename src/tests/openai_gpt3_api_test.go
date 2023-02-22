package openaigpt3bot_test

import (
	"os"
	"testing"

	gpt3 "github.com/klee1611/openai-gpt3-tg-bot"
)

func TestGPT3(t *testing.T) {
	envSetup()
	defer envTearDown()

	t.Run("TestGPT3NormalResponse", func(t *testing.T) {
		mockOpenAIServer := mockOpenAIServer()
		defer mockOpenAIServer.Close()

		gpt3API := gpt3.OpenAIGPT3API{
			BaseURL: mockOpenAIServer.URL,
			Token:   os.Getenv("OPENAI_API_KEY"),
		}

		res, err := gpt3API.GPT3("Test request message")
		if err != nil {
			t.Errorf("GPT3() error = %v", err)
		}
		if res != "Hello, world!" {
			t.Errorf("GPT3() = %v, want %v", res, "Hello, world!")
		}
	})
}
