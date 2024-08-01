package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gkwa/graybeauty/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		reader := strings.NewReader(path)

		englishTokenizer, err := core.NewEnglishTokenizer()
		if err != nil {
			return fmt.Errorf("NewEnglishTokenizer returned error: %v", err)
		}

		splitter := core.NewSentenceSplitter(englishTokenizer)

		var writer bytes.Buffer
		err = splitter.SplitSentences(reader, &writer)
		if err != nil {
			return errors.New("could ")
		}

		if err := os.WriteFile(path, writer.Bytes(), 0o600); err != nil {
			fmt.Println("Error:", err)
			return fmt.Errorf("error writing back to %s: %v", path, err)

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
