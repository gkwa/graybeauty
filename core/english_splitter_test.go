package core

import (
	"bytes"
	"strings"
	"testing"
)

func TestEnglishSplitter_SplitSentences(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "Simple sentences",
			input: `This is a test. Another sentence here.`,
			expected: `This is a test.

Another sentence here.

`,
		},
		{
			name:  "Complex sentence",
			input: `Mr. Smith went to Washington D.C. He had a meeting.`,
			expected: `Mr. Smith went to Washington D.C.

He had a meeting.

`,
		},
		{
			name: "Multiple lines",
			input: `Line one.
Line two.
Line three.`,
			expected: `Line one.

Line two.

Line three.

`,
		},
	}

	splitter, err := NewEnglishSplitter()
	if err != nil {
		t.Fatalf("Failed to create EnglishSplitter: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var writer bytes.Buffer

			err := splitter.SplitSentences(reader, &writer)
			if err != nil {
				t.Fatalf("SplitSentences returned an error: %v", err)
			}

			result := writer.String()
			if result != tc.expected {
				t.Errorf("Expected output:\n%s\nBut got:\n%s", tc.expected, result)
			}
		})
	}
}
