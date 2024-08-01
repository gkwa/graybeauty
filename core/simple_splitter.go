package core

import (
	"bufio"
	"io"
	"strings"
)

type SimpleSplitter struct{}

func NewSimpleSplitter() *SimpleSplitter {
	return &SimpleSplitter{}
}

func (s *SimpleSplitter) SplitSentences(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		sentences := strings.Split(line, ".")

		for i, sentence := range sentences {
			trimmedSentence := strings.TrimSpace(sentence)
			if trimmedSentence != "" {
				if i < len(sentences)-1 || strings.HasSuffix(line, ".") {
					trimmedSentence += "."
				}
				if _, err := w.Write([]byte(trimmedSentence + "\n\n")); err != nil {
					return err
				}
			}
		}
	}

	return scanner.Err()
}