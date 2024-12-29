package problems

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ParseConf struct {
	Dir                    string
	Section                string
	ReplacePackageWithMain bool
	ReplaceWithLiveLinks   bool
	Version                string
}

const (
	repoAddress = "https://github.com/spring1843/go-dsa/tree/"
	mainBranch  = "main"
)

var rehearsalRegex = regexp.MustCompile(`(?s)(## Rehearsal\n.*?)(\n##|\z)`)

func ParseSection(conf *ParseConf) (string, error) {
	content, err := os.ReadFile(filepath.Join(conf.Dir, conf.Section, "README.md"))
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	preparedContent, err := replaceRehearsal(string(content), conf)
	if err != nil {
		return "", fmt.Errorf("error replacing the rehearsal section in file: %w", err)
	}

	if conf.ReplaceWithLiveLinks {
		preparedContent = ReplaceWithLiveLinks(preparedContent, conf.Section, conf)
	}

	return preparedContent, nil
}

func replaceRehearsal(input string, conf *ParseConf) (string, error) {
	match := rehearsalRegex.FindStringSubmatch(input)
	if match == nil {
		return "", errors.New("no rehearsal section found")
	}

	rehearsalContent := match[1]
	newRehearsals, err := newRehearsalEntry(rehearsalContent)
	if err != nil {
		return "", fmt.Errorf("error parsing rehearsal entry: %w", err)
	}
	if len(newRehearsals) == 0 {
		log.Printf("no rehearsals found %s, %s", conf.Section, rehearsalContent)
	}
	replacement := fmt.Sprintf("## Rehearsal\n%s", stringRehearsalEntries(conf, newRehearsals))
	return rehearsalRegex.ReplaceAllString(input, replacement), nil
}

// ReplaceWithLiveLinks replaces the relative links in the input with links to the repo that work.
func ReplaceWithLiveLinks(input, section string, conf *ParseConf) string {
	repoBase := repoAddress + conf.Version
	if conf.Version == "" {
		repoBase = repoAddress + mainBranch
	}

	output := input
	if section != "" {
		output = strings.ReplaceAll(output, "](../", "]("+repoBase+"/")
	}

	return strings.ReplaceAll(output, "](./", "]("+repoBase+"/")
}
