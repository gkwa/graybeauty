package core

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/go-logr/logr/funcr"
)

func TestSplitSentences(t *testing.T) {
	var logBuf bytes.Buffer
	logger := funcr.New(func(prefix, args string) {
		logBuf.WriteString(prefix + args + "\n")
	}, funcr.Options{Verbosity: 1})

	englishTokenizer, err := NewEnglishTokenizer()
	if err != nil {
		t.Fatalf("Failed to create English tokenizer: %v", err)
	}

	splitter := NewSentenceSplitter(englishTokenizer, logger)

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single sentence",
			input:    "This is a single sentence.",
			expected: "This is a single sentence.\n\n",
		},
		{
			name:  "Multiple sentences",
			input: "This is the first sentence. This is the second sentence. And this is the third.",
			expected: `This is the first sentence.

This is the second sentence.

And this is the third.

`,
		},
		{
			name:  "Sentence with abbreviation",
			input: "Mr. Smith went to Washington D.C. He had a meeting.",
			expected: `Mr. Smith went to Washington D.C.

He had a meeting.

`,
		},
		{
			name:  "Multiple paragraphs",
			input: `Because this particular steamer has relatively small holes in the steamer tray, I cook the rice directly on the tray. The holes are small enough that very little rice will fall into the simmering water below. If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.`,
			expected: `Because this particular steamer has relatively small holes in the steamer tray, I cook the rice directly on the tray.

The holes are small enough that very little rice will fall into the simmering water below.

If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.

`,
		},
		{
			name: "Random bits added",
			input: `Because this particular steamer has

[[test this and that]]

relatively small holes in the steamer tray, I cook the rice directly on the tray. The holes are small enough that very little rice will fall into the simmering water below. If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.`,
			expected: `Because this particular steamer has

[[test this and that]]

relatively small holes in the steamer tray, I cook the rice directly on the tray.

The holes are small enough that very little rice will fall into the simmering water below.

If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.

`,
		},
		{
			name: "Random bits with newline squeezing",
			input: `Because this particular steamer has


[[test this and that]]









relatively small holes in the steamer tray, I cook the rice directly 










on the tray. The holes are small enough that very little rice will fall into the simmering water below. If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.`,
			expected: `Because this particular steamer has

[[test this and that]]

relatively small holes in the steamer tray, I cook the rice directly

on the tray.

The holes are small enough that very little rice will fall into the simmering water below.

If the holes on your steamer are any bigger, simply line the bottom of the steamer tray with a piece of cheesecloth before adding the rice.

`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var writer bytes.Buffer

			err := splitter.SplitSentences(reader, &writer)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			result := writer.String()
			if result != tc.expected {
				t.Errorf("Mismatch in split sentences:\n%s", diffStrings(tc.expected, result))
			}
		})
	}

	logOutput := logBuf.String()
	expectedLogs := []string{
		"Debug: Entering SplitSentences function",
		"Debug: Exiting SplitSentences function",
	}

	for _, expectedLog := range expectedLogs {
		if !strings.Contains(logOutput, expectedLog) {
			t.Errorf("Expected log output to contain '%s', but it didn't. Got: %s", expectedLog, logOutput)
		}
	}
}

func diffStrings(expected, actual string) string {
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")

	var diff strings.Builder
	diff.WriteString("Differences:\n")

	maxLines := len(expectedLines)
	if len(actualLines) > maxLines {
		maxLines = len(actualLines)
	}

	for i := 0; i < maxLines; i++ {
		var expectedLine, actualLine string
		if i < len(expectedLines) {
			expectedLine = expectedLines[i]
		}
		if i < len(actualLines) {
			actualLine = actualLines[i]
		}

		if expectedLine != actualLine {
			diff.WriteString(fmt.Sprintf("Line %d:\n", i+1))
			diff.WriteString(fmt.Sprintf("  Expected: %q\n", expectedLine))
			diff.WriteString(fmt.Sprintf("  Actual:   %q\n", actualLine))
		}
	}

	return diff.String()
}
