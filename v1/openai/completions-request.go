package openai

type CompletionsRequest struct {
	Echo             bool              `json:"echo,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Model            string            `json:"model,omitempty"`
	N                int               `json:"n,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	Prompt           string            `json:"prompt,omitempty"`
	Stop             string            `json:"stop,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	User             string            `json:"user,omitempty"`
}

func (c *CompletionsRequest) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}

func (c *CompletionsRequest) AddUserMessage(content string) {
	c.AddMessage(Message{
		Role:    "user",
		Content: content,
	})
}

func NewCompletions() *CompletionsRequest {
	return &CompletionsRequest{
		Model:       "gpt-4-1106-preview",
		N:           1,
		Temperature: 1,
		TopP:        1,
	}
}
