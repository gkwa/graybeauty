package core

type Tokenizer interface {
	Tokenize(text string) []string
}
