package openai

type Message struct {
	Content    string     `json:"content,omitempty"`
	Name       string     `json:"name,omitempty"`
	Role       string     `json:"role,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallId string     `json:"tool_call_id,omitempty"`
}

func (m *Message) IsValid() bool {
	return m.Role != "" && m.Content != ""
}

func (m *Message) SetContent(content string) *Message {
	if content != "" {
		m.Content = content
	}

	return m
}

func (m *Message) SetName(name string) *Message {
	if name != "" {
		m.Name = name
	}

	return m
}

func (m *Message) SetRole(role string) *Message {
	if role != "" {
		m.Role = role
	}

	return m
}

func NewMessage(content string) (newMessage *Message) {
	return &Message{
		Role:    "user",
		Content: content,
	}
}
