package cmd

import (
	"fmt"

	"github.com/gkwa/graybeauty/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Split sentences in a file",
	Args:  cobra.ExactArgs(1),
	Long:  `This command splits sentences in the specified file using the core package's sentence splitter.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		err := core.SplitSentencesInFile(path)
		if err != nil {
			return fmt.Errorf("error processing file: %v", err)
		}
		fmt.Printf("Successfully processed file: %s\n", path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
