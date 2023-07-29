package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	challengeCodePrefix = `// To get a random rehearsal challenge from go-dsa press the Run button on top right and browse the randomly picked URL.
package main

import (
	"fmt"
	"math/rand"
)

var rehearsals = `
	challengeCodeSuffix = `

func main() {
	fmt.Println(rehearsals[rand.Intn(len(rehearsals))])
}
`
	challengeURLPrefix = "https://github.com/spring1843/go-dsa/blob/master"
)

var randomChallengeCommand = &cobra.Command{
	Use:   "random-challenge",
	Short: "Generate code that randomly selects a rehearsal",
	Long:  "Discovers all challenges and outputs Go code that randomly selects one.",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error finding current directory: %s", err)
		}

		allChallenges := []string{}
		for _, section := range sections {
			matches, err := filepath.Glob(filepath.Join(dir, section, "*_test.go"))
			if err != nil {
				log.Fatalf("Error while globbing: %s", err)
			}

			for _, match := range matches {
				allChallenges = append(allChallenges, filepath.Join(challengeURLPrefix, section, filepath.Base(match)))
			}
		}
		fmt.Printf("%s%#v%s", challengeCodePrefix, allChallenges, challengeCodeSuffix)
	},
}
