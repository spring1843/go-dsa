package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Utility to work with go-dsa",
	Long:  `This is a CLI tool that is made to work with go-dsa.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.AddCommand(countRehearsalsCommand)
	rootCmd.AddCommand(exportMDCommand)
	rootCmd.AddCommand(randomChallengeCommand)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	rootCmd.AddCommand()
}
