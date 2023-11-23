package openai

type ToolCall struct {
	Function Function `json:"function,omitempty"`
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type,omitempty"`
}
