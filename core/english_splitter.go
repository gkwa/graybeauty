package core

import (
	"bufio"
	"io"
	"strings"
)

type EnglishSplitter struct {
	tokenizer Tokenizer
}

func NewEnglishSplitter() (*EnglishSplitter, error) {
	tokenizer, err := NewEnglishTokenizer()
	if err != nil {
		return nil, err
	}
	return &EnglishSplitter{tokenizer: tokenizer}, nil
}

func (s *EnglishSplitter) SplitSentences(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		sentences := s.tokenizer.Tokenize(line)

		for _, sentence := range sentences {
			trimmedSentence := strings.TrimSpace(sentence)
			if _, err := w.Write([]byte(trimmedSentence + "\n\n")); err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
