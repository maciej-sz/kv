package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a value to a key-value file",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("get called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
