package problems

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type rehearsalEntry struct {
	Name             string
	TestFileName     string
	SolutionFileName string
}

var rehearsalLinkPattern = regexp.MustCompile(`\* \[([^\]]+)\]\(\.\/([^\)]+)\), \[Solution\]\(\.\/([^\)]+)\)`)

func (r *rehearsalEntry) String() string {
	return fmt.Sprintf("### %s\n", r.Name)
}

/*
newRehearsalEntry parses the rehearsal section of the README.md file that looks like:

## Rehearsal

* [Some Name 1](./problem_test1.go), [Solution](./solution1.go)
* [Some Name 2](./problem_test2.go), [Solution](./solution2.go).
*/
func newRehearsalEntry(input string) ([]rehearsalEntry, error) {
	lines := strings.Split(input, "\n")
	entries := []rehearsalEntry{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "##") {
			continue // Skip empty lines and heading lines
		}
		matches := rehearsalLinkPattern.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, fmt.Errorf("invalid line format: %s, %d", line, len(matches))
		}

		entry := rehearsalEntry{
			Name:             matches[1],
			TestFileName:     matches[2],
			SolutionFileName: matches[3],
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func stringRehearsalEntries(conf *ParseConf, entries []rehearsalEntry) string {
	output := ""
	for _, entry := range entries {
		output += fmt.Sprintf("\n### %s\n", entry.Name)

		testFileContent, err := goFileFilters(filepath.Join(conf.Dir, conf.Section, entry.TestFileName), conf)
		if err != nil {
			output += fmt.Sprintf("Error reading test file: %s\n", err)
			continue
		}
		output += "\n```GO\n" + string(testFileContent) + "\n```\n"
	}

	output += "\n## Rehearsal Solutions\n"

	for _, entry := range entries {
		output += fmt.Sprintf("\n### %s\n", entry.Name)

		testFileContent, err := goFileFilters(filepath.Join(conf.Dir, conf.Section, entry.SolutionFileName), conf)
		if err != nil {
			output += fmt.Sprintf("Error reading test file: %s\n", err)
			continue
		}
		output += "\n```GO\n" + string(testFileContent) + "\n```\n"
	}

	return output
}

func goFileFilters(filePath string, conf *ParseConf) (string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading solution file [%s]: %s", filePath, err)
	}

	fileContentStr := string(fileContent)

	if conf.ReplacePackageWithMain {
		re := regexp.MustCompile(`(?m)^package ` + conf.Section + `$`)
		fileContentStr = re.ReplaceAllString(fileContentStr, "package main")
	}

	return fileContentStr, nil
}
