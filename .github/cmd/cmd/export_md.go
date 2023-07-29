package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var exportMDCommand = &cobra.Command{
	Use:   "export-md",
	Short: "Joins the entire project into one output",
	Long:  "Iteratively finds every README.md and outputs them",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error finding current directory: %s", err)
		}

		allFiles := []string{
			"README.md",
			"preface.md",
			"complexity.md",
		}
		for _, section := range sections {
			allFiles = append(allFiles, filepath.Join(section, "README.md"))
		}

		for _, file := range allFiles {
			content, err := os.ReadFile(filepath.Join(dir, file))
			if err != nil {
				log.Fatalf("Error reading file: %s", err)
			}
			fmt.Println(string(content))
		}
	},
}
