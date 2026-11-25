package problems

import (
	"errors"
	"fmt"
	"io/fs"
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

var errInvalidLine = errors.New("invalid line format")

var rehearsalLinkPattern = regexp.MustCompile(`[*-] \[([^\]]+)\]\(\.\/([^\)]+)\), \[Solution\]\(\.\/([^\)]+)\)`)

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
			return nil, fmt.Errorf("%w: %s, %d", errInvalidLine, line, len(matches))
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

func stringRehearsalEntries(conf *ParseConf, entries []rehearsalEntry) (string, error) {
	var output strings.Builder
	for _, entry := range entries {
		fmt.Fprintf(&output, "\n### %s\n", entry.Name)

		testFileContent, err := goFileFilters(entry.TestFileName, conf)
		if err != nil {
			return "", fmt.Errorf("error exporting test file for %q: %w", entry.Name, err)
		}
		output.WriteString("\n```GO\n" + testFileContent + "\n```\n")
	}

	output.WriteString("\n## Rehearsal Solutions\n")

	for _, entry := range entries {
		fmt.Fprintf(&output, "\n### %s\n", entry.Name)

		solutionFileContent, err := goFileFilters(entry.SolutionFileName, conf)
		if err != nil {
			return "", fmt.Errorf("error exporting solution file for %q: %w", entry.Name, err)
		}
		output.WriteString("\n```GO\n" + solutionFileContent + "\n```\n")
	}

	return output.String(), nil
}

func goFileFilters(fileName string, conf *ParseConf) (string, error) {
	fileContent, err := fs.ReadFile(os.DirFS(filepath.Join(conf.Dir, conf.Section)), fileName)
	if err != nil {
		return "", fmt.Errorf("error reading file [%s]: %w", fileName, err)
	}

	fileContentStr := string(fileContent)

	if conf.ReplacePackageWithMain {
		re := regexp.MustCompile(`(?m)^package ` + conf.Section + `$`)
		fileContentStr = re.ReplaceAllString(fileContentStr, "package main")
	}

	return fileContentStr, nil
}
