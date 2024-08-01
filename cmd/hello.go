package cmd

import (
	"fmt"

	"github.com/gkwa/graybeauty/core"
	"github.com/spf13/cobra"
)

var (
	writeFlag bool
	helloCmd  = &cobra.Command{
		Use:   "hello [file]",
		Short: "Split sentences in a file",
		Args:  cobra.ExactArgs(1),
		Long:  `This command splits sentences in the specified file using the core package's sentence splitter.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			result, err := core.ProcessFile(path, writeFlag)
			if err != nil {
				return fmt.Errorf("error processing file: %v", err)
			}
			fmt.Println(result)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolVarP(&writeFlag, "write", "w", false, "Write result back to file")
}
