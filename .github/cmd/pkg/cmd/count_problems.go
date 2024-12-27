package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spring1843/go-dsa/.github/cmd/pkg/problems"

	"github.com/spf13/cobra"
)

var countRehearsalsCommand = &cobra.Command{
	Use:   "count-rehearsals",
	Short: "Counts the number of rehearsals",
	Long:  "Iteratively counts the number of rehearsals for each section.",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error finding current directory: %s", err)
		}

		count := 0
		for _, section := range problems.OrderedSections {
			matches, err := filepath.Glob(filepath.Join(dir, section, "*_test.go"))
			if err != nil {
				log.Fatalf("Error while globbing: %s", err)
			}
			count += len(matches)
		}
		fmt.Println(count)
	},
}
