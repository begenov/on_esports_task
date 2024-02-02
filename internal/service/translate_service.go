package service

import (
	"cloud.google.com/go/translate"
)

type TranslationService interface {
	Translate(text, targetLanguage string) (string, error)
}

type GoogleTranslationService struct {
	client *translate.Client
}

func NewTranslateService(apiKey string) (*GoogleTranslationService, error) {
	return &GoogleTranslationService{}, nil
}

func (s *GoogleTranslationService) Translate(text, language string) (string, error) {
	return "", nil

}
