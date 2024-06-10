package cmd

import (
	"github.com/oxio/kv/pkg"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	const (
		defaultValFlag = "default"
	)
	var defaultVal *string

	cmd := &cobra.Command{
		Use:   "get <file> <key> [--default|-d value] ",
		Short: "Gets a value from the key-value file",
		RunE: func(cmd *cobra.Command, args []string) error {
			var val string
			var err error
			file := args[0]
			key := args[1]

			repo := pkg.NewKvRepo(file, &pkg.KvParser{})

			if cmd.Flag(defaultValFlag).Changed {
				val, err = repo.Find(key, defaultVal)
			} else {
				val, err = repo.Get(key)
			}

			if err != nil {
				return err
			}

			cmd.Print(val)

			return nil
		},
	}

	defaultVal = cmd.Flags().StringP(defaultValFlag, "d", "", "Default value")

	return cmd
}

func init() {
	rootCmd.AddCommand(newGetCmd())
}
