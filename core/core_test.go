package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcessFile(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	content := []byte(`This is a test. It has multiple sentences. How will it be split?`)
	err := os.WriteFile(testFile, content, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	testCases := []struct {
		name          string
		writeFlag     bool
		expectedOut   string
		expectedError bool
	}{
		{
			name:      "Process without write",
			writeFlag: false,
			expectedOut: `This is a test.

It has multiple sentences.

How will it be split?

`,
			expectedError: false,
		},
		{
			name:          "Process with write",
			writeFlag:     true,
			expectedOut:   "Successfully processed and updated file: " + testFile,
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ProcessFile(testFile, tc.writeFlag)

			if tc.expectedError && err == nil {
				t.Errorf("Expected an error, but got none")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != tc.expectedOut {
				t.Errorf("Expected output:\n%s\nBut got:\n%s", tc.expectedOut, result)
			}

			if tc.writeFlag {
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
		})
	}
}
