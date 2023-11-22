package openai

import (
	"io"
	"log"
	"os"
)

type TranscriptionRequest struct {
	File           []byte  `json:"file,omitempty"`
	Language       string  `json:"language,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
}

func (t *TranscriptionRequest) LoadFile(filename string) *TranscriptionRequest {
	content, err := os.Open(filename)

	if err != nil {
		log.Println(err)
		return t
	}

	file, err := io.ReadAll(content)

	if err != nil {
		log.Println(err)
		return t
	}

	return t.SetFile(file)
}

func (t *TranscriptionRequest) SetFile(file []byte) *TranscriptionRequest {
	if file != nil {
		t.File = file
	}

	return t
}

func (t *TranscriptionRequest) SetLanguage(language string) *TranscriptionRequest {
	if language != "" {
		t.Language = language
	}

	return t
}

func (t *TranscriptionRequest) SetModel(model string) *TranscriptionRequest {
	if model != "" {
		t.Model = model
	}

	return t
}

func (t *TranscriptionRequest) SetPrompt(prompt string) *TranscriptionRequest {
	if prompt != "" {
		t.Prompt = prompt
	}

	return t
}

func (t *TranscriptionRequest) SetResponseFormat(responseFormat string) *TranscriptionRequest {
	if responseFormat != "" {
		t.ResponseFormat = responseFormat
	}

	return t
}

func (t *TranscriptionRequest) SetTemperature(temperature float64) *TranscriptionRequest {
	if temperature > 0 {
		t.Temperature = temperature
	}

	return t
}

func NewTranscription() *TranscriptionRequest {
	return &TranscriptionRequest{
		ResponseFormat: "json",
		Language:       "en",
		Model:          "whisper-1",
		Temperature:    0,
	}
}
