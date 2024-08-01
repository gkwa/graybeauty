package core

import (
	"strings"

	"github.com/neurosnap/sentences"
	"github.com/neurosnap/sentences/english"
)

type EnglishTokenizer struct {
	tokenizer sentences.SentenceTokenizer
}

func NewEnglishTokenizer() (*EnglishTokenizer, error) {
	training, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		return nil, err
	}
	return &EnglishTokenizer{tokenizer: training}, nil
}

func (t *EnglishTokenizer) Tokenize(text string) []string {
	sentences := t.tokenizer.Tokenize(text)
	var result []string
	for _, s := range sentences {
		result = append(result, strings.TrimSpace(s.Text))
	}
	return result
}
