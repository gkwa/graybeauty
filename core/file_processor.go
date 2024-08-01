package core

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	splitter *SentenceSplitter
}

func NewFileProcessor() (*FileProcessor, error) {
	splitter, err := NewDefaultSentenceSplitter()
	if err != nil {
		return nil, fmt.Errorf("error creating sentence splitter: %v", err)
	}
	return &FileProcessor{splitter: splitter}, nil
}

func (fp *FileProcessor) ProcessFile(path string, writeFlag bool) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	processedContent, err := fp.splitter.Process(content)
	if err != nil {
		return "", fmt.Errorf("error processing content: %v", err)
	}

	if writeFlag {
		err = fp.writeProcessedContent(path, processedContent)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Successfully processed and updated file: %s", path), nil
	}

	return string(processedContent), nil
}

func (fp *FileProcessor) writeProcessedContent(path string, processedContent []byte) error {
	originalContent, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading original file: %v", err)
	}

	if bytes.Equal(originalContent, processedContent) {
		return fmt.Errorf("no changes made to file: %s", path)
	}

	tempDir := filepath.Dir(path)
	tempFile, err := os.CreateTemp(tempDir, "graybeauty-*")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath) // Clean up the temp file in case of failure

	_, err = tempFile.Write(processedContent)
	if err != nil {
		return fmt.Errorf("error writing to temp file: %v", err)
	}

	err = tempFile.Close()
	if err != nil {
		return fmt.Errorf("error closing temp file: %v", err)
	}

	err = os.Rename(tempFilePath, path)
	if err != nil {
		return fmt.Errorf("error replacing original file with processed content: %v", err)
	}

	return nil
}
