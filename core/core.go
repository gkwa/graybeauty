package core

import (
	"fmt"
)

func ProcessFile(path string, writeFlag bool) (string, error) {
	splitter, err := NewEnglishSplitter()
	if err != nil {
		return "", fmt.Errorf("error creating English splitter: %v", err)
	}

	processor := NewFileProcessor(splitter)
	return processor.ProcessFile(path, writeFlag)
}
