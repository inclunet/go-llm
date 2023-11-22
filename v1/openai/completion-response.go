package openai

type CompletionsResponse struct {
	Created int    `json:"created,omitempty"`
	Error   *Error `json:"error,omitempty"`
	ID      string `json:"id,omitempty"`
	Model   string `json:"model,omitempty"`
	Object  string `json:"object,omitempty"`

	Choices []struct {
		Message      Message     `json:"message"`
		Text         string      `json:"text,omitempty"`
		Index        int         `json:"index,omitempty"`
		Logprobs     interface{} `json:"logprobs,omitempty"`
		FinishReason string      `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`

	Usage struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`
}
