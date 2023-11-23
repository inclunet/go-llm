package openai

import (
	"errors"
	"io/ioutil"
)

type Speech struct {
	Input          string `json:"input,omitempty"`
	Model          string `json:"model,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	Speed          int    `json:"speed,omitempty"`
	Voice          string `json:"voice,omitempty"`
}

func (s *Speech) SaveToFile(content []byte, path string) error {
	if path == "" || content == nil {
		return errors.New("path or content is empty")
	}

	return ioutil.WriteFile(path+"."+s.ResponseFormat, content, 0644)
}

func (s *Speech) SetInput(input string) *Speech {
	if input != "" {
		s.Input = input
	}

	return s
}

func (s *Speech) SetModel(model string) *Speech {
	if model != "" {
		s.Model = model
	}

	return s
}

func (s *Speech) SetSpeed(speed int) *Speech {
	if speed > 0 {
		s.Speed = speed
	}

	return s
}

func (s *Speech) SetResponseFormat(responseFormat string) *Speech {
	if responseFormat != "" {
		s.ResponseFormat = responseFormat
	}

	return s
}

func (s *Speech) SetVoice(voice string) *Speech {
	if voice != "" {
		s.Voice = voice
	}

	return s
}

func NewSpeech() *Speech {
	return &Speech{
		Model:          "tts-1-hd",
		ResponseFormat: "mp3",
		Speed:          1,
		Voice:          "alloy",
	}
}
