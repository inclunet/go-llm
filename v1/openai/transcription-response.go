package openai

import (
	"errors"
	"io/ioutil"
)

type TranscriptionResponse struct {
	Text string `json:"text,omitempty"`
}

func (t *TranscriptionResponse) GetText() string {
	return t.Text
}

func (t *TranscriptionResponse) SaveToFile(filename string) error {
	if filename == "" {
		return errors.New("filename is empty")
	}

	content := []byte(t.Text)

	return ioutil.WriteFile(filename, content, 0644)
}
