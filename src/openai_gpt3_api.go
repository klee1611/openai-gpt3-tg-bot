package openaigpt3bot

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type OpenAIGPT3API struct {
	BaseURL string
	Token   string
}

type OpenAIGPT3Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func (c *OpenAIGPT3API) GPT3(message string) (res string, err error) {
	reqBody := map[string]interface{}{
		"model":             "gpt-3.5-turbo",
		"prompt":            message,
		"max_tokens":        256,
		"temperature":       0.9,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0.6,
	}

	log.Info("Sending request to OpenAI GPT3...")
	resp, err := client.R().
		SetBody(reqBody).
		SetAuthToken(c.Token).
		SetHeader("Content-Type", "application/json").
		SetResult(map[string]interface{}{}).
		Post(c.BaseURL)
	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Errorf("Error response from OpenAI GPT3 API: %v", resp)
		return "", err
	}
	log.Info("Got response from OpenAI GPT3")

	var result OpenAIGPT3Response
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Error(err)
		return
	}

	return result.Choices[0].Text, err
}
