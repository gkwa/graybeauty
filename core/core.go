package core

import (
	"bufio"
	"io"
	"strings"

	"github.com/go-logr/logr"
	"github.com/neurosnap/sentences/english"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}

func SplitSentences(logger logr.Logger, r io.Reader, w io.Writer) error {
	logger.V(1).Info("Debug: Entering SplitSentences function")

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		logger.Error(err, "Failed to create tokenizer")
		return err
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		sentenceSegments := tokenizer.Tokenize(line)

		for _, s := range sentenceSegments {
			_, err := w.Write([]byte(strings.TrimSpace(s.Text) + "\n\n"))
			if err != nil {
				logger.Error(err, "Failed to write sentence")
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		logger.Error(err, "Failed to read input")
		return err
	}

	logger.V(1).Info("Debug: Exiting SplitSentences function")
	return nil
}
