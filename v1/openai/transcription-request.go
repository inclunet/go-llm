package openai

import (
	"io"
	"log"
	"os"
)

// TranscriptionRequest is the request for the transcription endpoint
// https://beta.openai.com/docs/api-reference/transcriptions/create
type TranscriptionRequest struct {
	File           []byte  `json:"file,omitempty"`
	Language       string  `json:"language,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
}

// LoadFile loads a file from the filesystem and sets it as the file for the request
// The file is the audio file to be transcribed
// https://beta.openai.com/docs/api-reference/transcriptions/create#file
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

// SetFile sets the file for the request
// The file is the audio file to be transcribed
// https://beta.openai.com/docs/api-reference/transcriptions/create#file
func (t *TranscriptionRequest) SetFile(file []byte) *TranscriptionRequest {
	if file != nil {
		t.File = file
	}

	return t
}

// SetLanguage sets the language for the request
// The language is the language of the audio file
// https://beta.openai.com/docs/api-reference/transcriptions/create#language
func (t *TranscriptionRequest) SetLanguage(language string) *TranscriptionRequest {
	if language != "" {
		t.Language = language
	}

	return t
}

// SetModel sets the model for the request
// The model is the model to use for transcription
// https://beta.openai.com/docs/api-reference/transcriptions/create#model
func (t *TranscriptionRequest) SetModel(model string) *TranscriptionRequest {
	if model != "" {
		t.Model = model
	}

	return t
}

// SetPrompt sets the prompt for the request
// The prompt is the prompt to use for transcription
// https://beta.openai.com/docs/api-reference/transcriptions/create#prompt
func (t *TranscriptionRequest) SetPrompt(prompt string) *TranscriptionRequest {
	if prompt != "" {
		t.Prompt = prompt
	}

	return t
}

// SetResponseFormat sets the response format for the request
// The response format is the format of the response
// https://beta.openai.com/docs/api-reference/transcriptions/create#response_format
func (t *TranscriptionRequest) SetResponseFormat(responseFormat string) *TranscriptionRequest {
	if responseFormat != "" {
		t.ResponseFormat = responseFormat
	}

	return t
}

// SetTemperature sets the temperature for the request
// The temperature is the temperature for the response
// https://beta.openai.com/docs/api-reference/transcriptions/create#temperature
func (t *TranscriptionRequest) SetTemperature(temperature float64) *TranscriptionRequest {
	if temperature > 0 {
		t.Temperature = temperature
	}

	return t
}

// NewTranscription returns a new TranscriptionRequest with default values
func NewTranscription() *TranscriptionRequest {
	return &TranscriptionRequest{
		ResponseFormat: "json",
		Language:       "en",
		Model:          "whisper-1",
		Temperature:    0,
	}
}
