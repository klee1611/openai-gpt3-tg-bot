package openaichatbot

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type OpenAIAPI struct {
	BaseURL string
	Token   string
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func (c *OpenAIAPI) OpenAI(message string) (res string, err error) {
	reqBody := map[string]interface{}{
		"model":             "text-davinci-003",
		"prompt":            message,
		"max_tokens":        256,
		"temperature":       0.7,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	log.Info("Sending request to openai")
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
		log.Errorf("Error response from openAI API: %v", resp)
		return "", err
	}
	log.Info("Got response from openai")

	var result OpenAIResponse
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Error(err)
		return
	}

	return result.Choices[0].Text, err
}
