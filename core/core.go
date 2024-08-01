package core

import (
	"fmt"
)

func ProcessFile(path string, writeFlag bool) (string, error) {
	processor, err := NewFileProcessor()
	if err != nil {
		return "", fmt.Errorf("error creating file processor: %v", err)
	}
	return processor.ProcessFile(path, writeFlag)
}
