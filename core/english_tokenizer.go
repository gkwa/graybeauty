package core

import (
	"strings"

	"github.com/neurosnap/sentences/english"
)

type EnglishTokenizer struct{}

func NewEnglishTokenizer() (*EnglishTokenizer, error) {
	_, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		return nil, err
	}
	return &EnglishTokenizer{}, nil
}

func (t *EnglishTokenizer) Tokenize(text string) []string {
	tokenizer, _ := english.NewSentenceTokenizer(nil)
	sentences := tokenizer.Tokenize(text)
	var result []string
	for _, s := range sentences {
		result = append(result, strings.TrimSpace(s.Text))
	}
	return result
}
