package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kv",
	Short: "Simple Key-Value storage tool",
}

func Execute() {
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	err := rootCmd.Execute()
	if err != nil {
		errorLogger.Println(err)
		os.Exit(1)
	}
}
