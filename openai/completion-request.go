package openai

import (
	"encoding/json"
	"io"
	"os"
)

// CompletionRequest is the request for the completion endpoint
// https://beta.openai.com/docs/api-reference/completions/create
type CompletionRequest struct {
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Model            string            `json:"model,omitempty"`
	N                int               `json:"n,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	ResponseFormat   string            `json:"response_format,omitempty"`
	Seed             float64           `json:"seed,omitempty"`
	Stop             string            `json:"stop,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	User             string            `json:"user,omitempty"`
}

// AddMessage adds a message to the request
// https://beta.openai.com/docs/api-reference/completions/create#messages
func (c *CompletionRequest) AddMessage(message Message) *CompletionRequest {
	if !message.IsValid() {
		c.Messages = append(c.Messages, message)
	}

	return c
}

// SaveToFile saves the request to a file
func (c *CompletionRequest) SaveToFile(filename string) error {
	file, err := json.MarshalIndent(c, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}

// SetFrequencyPenalty sets the frequency penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#frequency_penalty
func (c *CompletionRequest) SetFrequencyPenalty(frequencyPenalty float64) *CompletionRequest {
	if frequencyPenalty > 0 {
		c.FrequencyPenalty = frequencyPenalty
	}

	return c
}

// SetLogitBias sets the logit bias for the request
// https://beta.openai.com/docs/api-reference/completions/create#logit_bias
func (c *CompletionRequest) SetLogitBias(logitBias map[string]string) *CompletionRequest {
	if logitBias != nil {
		c.LogitBias = logitBias
	}

	return c
}

// SetMaxTokens sets the max tokens for the request
// https://beta.openai.com/docs/api-reference/completions/create#max_tokens
func (c *CompletionRequest) SetMaxTokens(maxTokens int) *CompletionRequest {
	if maxTokens > 0 {
		c.MaxTokens = maxTokens
	}

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetModel(model string) *CompletionRequest {
	if model != "" {
		c.Model = model
	}

	return c
}

// SetN sets the n for the request
// https://beta.openai.com/docs/api-reference/completions/create#n
func (c *CompletionRequest) SetN(n int) *CompletionRequest {
	if n > 0 {
		c.N = n
	}

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetTemperature(temperature float64) *CompletionRequest {
	if temperature > 0 {
		c.Temperature = temperature
	}

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetTopP(topP float64) *CompletionRequest {
	if topP > 0 {
		c.TopP = topP
	}

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetStream(stream bool) *CompletionRequest {
	c.Stream = stream

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetStop(stop string) *CompletionRequest {
	if stop != "" {
		c.Stop = stop
	}

	return c
}

// SetPresencePenalty sets the presence penalty for the request
// https://beta.openai.com/docs/api-reference/completions/create#presence_penalty
func (c *CompletionRequest) SetUser(user string) *CompletionRequest {
	if user != "" {
		c.User = user
	}

	return c
}

// LoadCompletionFromFile loads a CompletionRequest from a file
func LoadCompletionFromFile(filename string) (*CompletionRequest, error) {
	content, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	file, err := io.ReadAll(content)

	if err != nil {
		return nil, err
	}

	completion := NewCompletions()

	err = json.Unmarshal(file, completion)

	return completion, err
}

// NewCompletions returns a new CompletionRequest
func NewCompletions() *CompletionRequest {
	return &CompletionRequest{
		Model:       "gpt-4-1106-preview",
		N:           1,
		Temperature: 1,
		TopP:        1,
	}
}
