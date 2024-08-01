package core

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSplitSentences(t *testing.T) {
	englishTokenizer, err := NewEnglishTokenizer()
	if err != nil {
		t.Fatalf("Failed to create English tokenizer: %v", err)
	}

	splitter := NewSentenceSplitter(englishTokenizer)

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "Single sentence",
			input: `This is a single sentence.`,
			expected: `This is a single sentence.

`,
		},
		{
			name:  "Multiple sentences",
			input: `This is the first sentence. This is the second sentence. And this is the third.`,
			expected: `This is the first sentence.

This is the second sentence.

And this is the third.

`,
		},
		{
			name:  "Sentence with abbreviation",
			input: `Mr. Smith went to Washington D.C. He had a meeting.`,
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
		{
			name: "Markdown embedded",
			input: `This is some markdown:

- item 1
- item 2
`,
			expected: `This is some markdown:

- item 1

- item 2

`,
		},
		{
			name: "Markdown embedded",
			input: `## Overview

A [Kubernetes operator](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000001OMH6AAO) is a method of packaging, deploying, and managing a Kubernetes application. A Kubernetes application is both deployed on [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes) and managed using the Kubernetes API (application programming interface) and kubectl tooling.

A Kubernetes operator is an application-specific controller that extends the functionality of the Kubernetes API to create, configure, and manage instances of complex applications on behalf of a Kubernetes user.

`,
			expected: `## Overview

A [Kubernetes operator](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000001OMH6AAO) is a method of packaging, deploying, and managing a Kubernetes application.

A Kubernetes application is both deployed on [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes) and managed using the Kubernetes API (application programming interface) and kubectl tooling.

A Kubernetes operator is an application-specific controller that extends the functionality of the Kubernetes API to create, configure, and manage instances of complex applications on behalf of a Kubernetes user.

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
