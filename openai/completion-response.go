package openai

// CompletionResponse is the response from the Completion API.
// See: https://beta.openai.com/docs/api-reference/completions/create
type CompletionResponse struct {
	Choices []Choice `json:"choices,omitempty"`
	Created int      `json:"created,omitempty"`
	Error   *Error   `json:"error,omitempty"`
	ID      string   `json:"id,omitempty"`
	Model   string   `json:"model,omitempty"`
	Object  string   `json:"object,omitempty"`
	Usage   Usage    `json:"usage,omitempty"`
}

func NewCompletionResponse() (newCompletionResponse *CompletionResponse) {
	return &CompletionResponse{}
}
