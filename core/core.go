package core

import (
	"bufio"
	"io"
	"strings"

	"github.com/go-logr/logr"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}

type Tokenizer interface {
	Tokenize(text string) []string
}

type SentenceSplitter struct {
	tokenizer Tokenizer
	logger    logr.Logger
}

func NewSentenceSplitter(tokenizer Tokenizer, logger logr.Logger) *SentenceSplitter {
	return &SentenceSplitter{
		tokenizer: tokenizer,
		logger:    logger,
	}
}

func (s *SentenceSplitter) SplitSentences(r io.Reader, w io.Writer) error {
	s.logger.V(1).Info("Debug: Entering SplitSentences function")

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		sentences := s.tokenizer.Tokenize(line)

		for _, sentence := range sentences {
			trimmedSentence := strings.TrimSpace(sentence)
			if _, err := w.Write([]byte(trimmedSentence + "\n\n")); err != nil {
				s.logger.Error(err, "Failed to write sentence")
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		s.logger.Error(err, "Failed to read input")
		return err
	}

	s.logger.V(1).Info("Debug: Exiting SplitSentences function")
	return nil
}
