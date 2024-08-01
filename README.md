# graybeauty
Purpose: A flexible sentence splitting tool for various text processing needs.

## Example Usage
```bash
# Process a file and print the result
graybeauty hello myfile.txt

# Process a file and write the changes back to the file
graybeauty hello -w myfile.txt
```

## Install graybeauty
On macOS/Linux:
```bash
brew install gkwa/homebrew-tools/graybeauty
```

On Windows:
```powershell
TBD
```

## Creating Custom Splitters

graybeauty uses the Strategy Pattern to allow for flexible sentence splitting algorithms. Here are a few ways to create and use custom splitters:

1. Implement the SplitterStrategy interface:
```go
type SplitterStrategy interface {
    SplitSentences(r io.Reader, w io.Writer) error
}
```

2. Create a new splitter:
```go
type CustomSplitter struct{}

func (s *CustomSplitter) SplitSentences(r io.Reader, w io.Writer) error {
    // Implement your custom splitting logic here
}
```

3. Use the custom splitter:
```go
customSplitter := &CustomSplitter{}
processor := NewFileProcessor(customSplitter)
result, err := processor.ProcessFile("path/to/file", writeFlag)
```

### Examples of Custom Splitters

1. RegexSplitter: Use regular expressions for splitting.
2. MLSplitter: Implement a machine learning model for more accurate sentence detection.
3. LanguageSpecificSplitter: Create splitters tailored for different languages.

## Extending graybeauty

To add a new splitting strategy:

1. Create a new file in the `core` package (e.g., `custom_splitter.go`).
2. Implement the `SplitterStrategy` interface.
3. Update `core.go` to allow using your new splitter.

Example:
```go
func ProcessFileWithCustomSplitter(path string, writeFlag bool) (string, error) {
    splitter := NewCustomSplitter()
    processor := NewFileProcessor(splitter)
    return processor.ProcessFile(path, writeFlag)
}
```

By following this pattern, graybeauty can be easily extended to support various text processing needs without modifying existing code.
```

This README now provides:

1. A brief explanation of the tool's purpose
2. Example usage of the CLI
3. Installation instructions
4. An overview of how to create custom splitters using the Strategy Pattern
5. Examples of potential custom splitters
6. Instructions on how to extend graybeauty with new splitting strategies

This documentation should help users understand how to use graybeauty and how to extend it for their specific needs.
