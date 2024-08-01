package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Tokenizer interface {
	Tokenize(text string) []string
}

type SentenceSplitter struct {
	tokenizer Tokenizer
}

func NewSentenceSplitter(tokenizer Tokenizer) *SentenceSplitter {
	return &SentenceSplitter{
		tokenizer: tokenizer,
	}
}

func ProcessFile(path string, writeFlag bool) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	englishTokenizer, err := NewEnglishTokenizer()
	if err != nil {
		return "", fmt.Errorf("error creating tokenizer: %v", err)
	}

	splitter := NewSentenceSplitter(englishTokenizer)
	var buf bytes.Buffer

	err = splitter.SplitSentences(bytes.NewReader(content), &buf)
	if err != nil {
		return "", fmt.Errorf("error splitting sentences: %v", err)
	}

	if writeFlag {
		if !bytes.Equal(content, buf.Bytes()) {
			err = os.WriteFile(path, buf.Bytes(), 0o644)
			if err != nil {
				return "", fmt.Errorf("error writing file: %v", err)
			}
			return fmt.Sprintf("Successfully processed and updated file: %s", path), nil
		}
		return fmt.Sprintf("No changes made to file: %s", path), nil
	}

	return buf.String(), nil
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
