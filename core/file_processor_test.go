package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileProcessor_ProcessFile_EnglishSplitter_WithoutWrite(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	content := []byte(`This is a test. It has multiple sentences. How will it be split?`)
	err := os.WriteFile(testFile, content, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	englishSplitter, err := NewEnglishSplitter()
	if err != nil {
		t.Fatalf("Failed to create EnglishSplitter: %v", err)
	}

	processor := NewFileProcessor(englishSplitter)
	result, err := processor.ProcessFile(testFile, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOut := `This is a test.

It has multiple sentences.

How will it be split?

`
	if result != expectedOut {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOut, result)
	}
}

func TestFileProcessor_ProcessFile_SimpleSplitter_WithWrite(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	content := []byte(`This is a test. It has multiple sentences. How will it be split?`)
	err := os.WriteFile(testFile, content, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	simpleSplitter := NewSimpleSplitter()
	processor := NewFileProcessor(simpleSplitter)
	result, err := processor.ProcessFile(testFile, true)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOut := "Successfully processed and updated file: " + testFile
	if result != expectedOut {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOut, result)
	}

	// Check if file was actually updated
	updatedContent, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}
	expectedContent := `This is a test.

It has multiple sentences.

How will it be split?

`
	if string(updatedContent) != expectedContent {
		t.Errorf("File content not updated as expected. Got:\n%s\nExpected:\n%s", string(updatedContent), expectedContent)
	}
}
