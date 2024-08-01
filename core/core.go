package core

import (
	"fmt"
	"os"
)

func ProcessFile(path string, writeFlag bool) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	splitter, err := NewDefaultSentenceSplitter()
	if err != nil {
		return "", fmt.Errorf("error creating sentence splitter: %v", err)
	}

	processedContent, err := splitter.Process(content)
	if err != nil {
		return "", fmt.Errorf("error processing content: %v", err)
	}

	if writeFlag {
		if string(content) != string(processedContent) {
			err = os.WriteFile(path, processedContent, 0)
			if err != nil {
				return "", fmt.Errorf("error writing file: %v", err)
			}
			return fmt.Sprintf("Successfully processed and updated file: %s", path), nil
		}
		return fmt.Sprintf("No changes made to file: %s", path), nil
	}

	return string(processedContent), nil
}
