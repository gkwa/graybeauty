package core

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type SentenceSplitter struct {
	tokenizer Tokenizer
}

func NewSentenceSplitter(tokenizer Tokenizer) *SentenceSplitter {
	return &SentenceSplitter{
		tokenizer: tokenizer,
	}
}

func NewDefaultSentenceSplitter() (*SentenceSplitter, error) {
	tokenizer, err := NewEnglishTokenizer()
	if err != nil {
		return nil, err
	}
	return NewSentenceSplitter(tokenizer), nil
}

func (s *SentenceSplitter) Process(content []byte) ([]byte, error) {
	reader := bytes.NewReader(content)
	var writer bytes.Buffer
	err := s.SplitSentences(reader, &writer)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func (s *SentenceSplitter) SplitSentences(r io.Reader, w io.Writer) error {
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
