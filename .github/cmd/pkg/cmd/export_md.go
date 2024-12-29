package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spring1843/go-dsa/.github/cmd/pkg/problems"

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

		nonSectionMDFiles := []string{
			"README.md",
			"preface.md",
			"complexity.md",
		}

		conf := &problems.ParseConf{
			Dir:                    dir,
			ReplacePackageWithMain: cmd.Flag(replacePackageWithMainFlag).Value.String() == "true",
			ReplaceWithLiveLinks:   cmd.Flag(replaceWithLiveLinksFlag).Value.String() == "true",
			Version:                cmd.Flag(versionFlag).Value.String(),
		}

		for _, file := range nonSectionMDFiles {
			content, err := os.ReadFile(filepath.Join(dir, file))
			if err != nil {
				log.Fatalf("Error reading file: %s", err)
			}

			contentStr := string(content)

			if conf.ReplaceWithLiveLinks {
				contentStr = problems.ReplaceWithLiveLinks(contentStr, "", conf)
			}

			fmt.Println(string(contentStr))
		}

		for _, section := range problems.OrderedSections {
			conf.Section = section

			parsedSection, err := problems.ParseSection(conf)
			if err != nil {
				log.Fatalf("Error parsing section: %s", err)
			}
			fmt.Println(parsedSection)
		}
	},
}
