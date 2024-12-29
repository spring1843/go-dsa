package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	replacePackageWithMainFlag = "replace-packages-with-main"
	replaceWithLiveLinksFlag   = "replace-with-live-links"
	versionFlag                = "version"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Utility to work with go-dsa",
	Long:  `This is a CLI tool that is made to work with go-dsa.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.AddCommand(countRehearsalsCommand)

	exportMDCommand.Flags().Bool(replacePackageWithMainFlag, true, "renames all packages in all Go files to use the main package")
	exportMDCommand.Flags().Bool(replaceWithLiveLinksFlag, true, "replace all links to the live version of the repository")
	exportMDCommand.Flags().String(versionFlag, "", "version of the release")

	rootCmd.AddCommand(exportMDCommand)
	rootCmd.AddCommand(randomChallengeCommand)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	rootCmd.AddCommand()
}
