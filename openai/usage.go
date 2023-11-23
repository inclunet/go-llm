package openai

// Usage is a struct that represents the usage of the API.
// See: https://beta.openai.com/docs/api-reference/usage/create
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
