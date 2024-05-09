package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a value from the key-value file",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("get called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
