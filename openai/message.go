package openai

type Message struct {
	Content    string     `json:"content,omitempty"`
	Name       string     `json:"name,omitempty"`
	Role       string     `json:"role,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallId string     `json:"tool_call_id,omitempty"`
}

// IsValid checks if the message is valid
func (m *Message) IsValid() bool {
	return m.Role != "" && m.Content != ""
}

// SetContent sets the content of the message
// The content is the text of the message
// https://beta.openai.com/docs/api-reference/messages/create#content
func (m *Message) SetContent(content string) *Message {
	if content != "" {
		m.Content = content
	}

	return m
}

// SetName sets the name of the message
// The name is the name of the user
// https://beta.openai.com/docs/api-reference/messages/create#name
func (m *Message) SetName(name string) *Message {
	if name != "" {
		m.Name = name
	}

	return m
}

// SetRole sets the role of the message
// The role is the role of the user
// https://beta.openai.com/docs/api-reference/messages/create#role
func (m *Message) SetRole(role string) *Message {
	if role != "" {
		m.Role = role
	}

	return m
}

// NewMessage creates a new message
// By default, the role is "user"
// https://beta.openai.com/docs/api-reference/messages/create#content
func NewMessage(content string) (newMessage Message) {
	return Message{
		Role:    "user",
		Content: content,
	}
}
