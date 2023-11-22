package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey       string
	organization string
	BaseURL      string
	UserAgent    string
}

func (c *Client) GetUrl(path string) string {
	return c.BaseURL + path
}

func (c *Client) getUrlWithParams(path string, params any) string {
	if params == nil {
		return c.GetUrl(path)
	}

	values := url.Values{}

	if paramList, ok := params.(map[string]string); ok {

		for key, value := range paramList {
			values.Add(key, value)
		}
	}

	queryString := values.Encode()

	if queryString != "" {
		queryString = "?" + queryString
	}

	return c.GetUrl(path) + queryString
}

func (c *Client) Get(path string, params any) (response []byte, err error) {
	resp, err := c.request(http.MethodGet, c.getUrlWithParams(path, params), nil)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)

	return response, err
}

func (c *Client) Post(path string, payload any) (response []byte, err error) {
	response = make([]byte, 0)

	jsonInput, err := json.Marshal(payload)

	if err != nil {
		return response, err
	}

	resp, err := c.request(http.MethodPost, c.GetUrl(path), bytes.NewReader(jsonInput))

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)

	return response, err
}

func (c *Client) request(method, url string, body io.Reader) (response *http.Response, err error) {
	newRequest, err := http.NewRequest(method, url, body)

	if err != nil {
		return response, err
	}

	newRequest.Header.Set("Authorization", "Bearer "+c.apiKey)
	newRequest.Header.Set("User-Agent", c.UserAgent)
	newRequest.Header.Set("Content-Type", "application/json")

	if c.organization != "" {
		newRequest.Header.Set("OpenAI-Organization", c.organization)
	}

	client := &http.Client{}
	response, err = client.Do(newRequest)

	return response, err
}

func (c *Client) Completions(completions *CompletionRequest) (response *CompletionsResponse, err error) {
	response = &CompletionsResponse{}

	jsonResponse, err := c.Post("/chat/completions", completions)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(jsonResponse, response)

	if err != nil {
		return response, err
	}

	if response.Error != nil {
		if err := response.Error.Error(); err != "" {
			return response, errors.New(err)
		}
	}

	for _, choice := range response.Choices {
		if choice.Message.Role == "assistant" {
			completions.AddMessage(choice.Message)
		}
	}

	return response, err
}

func (c *Client) Speech(speech *Speech) (response []byte, err error) {
	return c.Post("/audio/speech", speech)
}

func NewClient(apiKey, organizationID string) (newClient *Client) {
	return &Client{
		apiKey:       apiKey,
		organization: organizationID,
		BaseURL:      "https://api.openai.com/v1",
		UserAgent:    "go-openai",
	}
}

func (c *Client) Transcriptions(transcription *TranscriptionRequest) (response *TranscriptionResponse, err error) {
	response = &TranscriptionResponse{}

	jsonResponse, err := c.Post("/audio/transcriptions", transcription)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(jsonResponse, response)

	if err != nil {
		return response, err
	}

	return response, err
}

func (c *Client) Translations(transcription *TranscriptionRequest) (response *TranscriptionResponse, err error) {
	response = &TranscriptionResponse{}

	jsonResponse, err := c.Post("/audio/translations", transcription)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(jsonResponse, response)

	if err != nil {
		return response, err
	}

	return response, err
}
