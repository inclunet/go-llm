package openai

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

func (c *CompletionRequest) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}

func (c *CompletionRequest) AddUserMessage(content string) {
	c.AddMessage(Message{
		Role:    "user",
		Content: content,
	})
}


func (c *CompletionRequest) SetFrequencyPenalty(frequencyPenalty float64) *CompletionRequest {
	if frequencyPenalty > 0 {
		c.FrequencyPenalty = frequencyPenalty
	}

	return c
}

func (c *CompletionRequest) SetLogitBias(logitBias map[string]string) *CompletionRequest {
	if logitBias != nil {
		c.LogitBias = logitBias
	}

	return c
}

func (c *CompletionRequest) SetMaxTokens(maxTokens int) *CompletionRequest {
	if maxTokens > 0 {
		c.MaxTokens = maxTokens
	}

	return c
}

func (c *CompletionRequest) SetModel(model string) *CompletionRequest {
	if model != "" {
		c.Model = model
	}

	return c
}

func (c *CompletionRequest) SetN(n int) *CompletionRequest {
	if n > 0 {
		c.N = n
	}

	func (c *CompletionRequest) SetTemperature(temperature float64) *CompletionRequest {
		if temperature > 0 {
			c.Temperature = temperature
		}

	return c
}

func (c *CompletionRequest) SetTopP(topP float64) *CompletionRequest {
	if topP > 0 {
		c.TopP = topP
	}

	return c
}

func (c *CompletionRequest) SetStream(stream bool) *CompletionRequest {
	c.Stream = stream

	return c
}

func (c *CompletionRequest) SetStop(stop string) *CompletionRequest {
	if stop != "" {
		c.Stop = stop
	}

	return c
}

func (c *CompletionRequest) SetUser(user string) *CompletionRequest {
	if user != "" {
		c.User = user
	}

	return c
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
