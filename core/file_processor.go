package core

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	splitter SplitterStrategy
}

func NewFileProcessor(splitter SplitterStrategy) *FileProcessor {
	return &FileProcessor{splitter: splitter}
}

func (fp *FileProcessor) ProcessFile(path string, writeFlag bool) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	var buf bytes.Buffer
	err = fp.splitter.SplitSentences(bytes.NewReader(content), &buf)
	if err != nil {
		return "", fmt.Errorf("error processing content: %v", err)
	}

	processedContent := buf.Bytes()

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
