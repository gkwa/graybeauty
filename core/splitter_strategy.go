package core

import "io"

// SplitterStrategy defines the interface for sentence splitting algorithms
type SplitterStrategy interface {
	SplitSentences(r io.Reader, w io.Writer) error
}
