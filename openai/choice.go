package openai

// Choice is a single choice from the completion.
// See: https://beta.openai.com/docs/api-reference/completions/create
type Choice struct {
	FinishReason string      `json:"finish_reason,omitempty"`
	Index        int         `json:"index,omitempty"`
	Logprobs     interface{} `json:"logprobs,omitempty"`
	Message      Message     `json:"message"`
	Text         string      `json:"text,omitempty"`
}
