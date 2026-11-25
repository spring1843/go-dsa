package problems

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseConf configures how sections are parsed.
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

var errNoRehearsalSection = errors.New("no rehearsal section found")

var rehearsalRegex = regexp.MustCompile(`(?s)(## Rehearsal\n.*?)(\n##|\z)`)

// ParseSection returns the README.md of a section with its rehearsal links replaced with code.
func ParseSection(conf *ParseConf) (string, error) {
	content, err := fs.ReadFile(os.DirFS(filepath.Join(conf.Dir, conf.Section)), "README.md")
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
		return "", errNoRehearsalSection
	}

	rehearsalContent := match[1]
	newRehearsals, err := newRehearsalEntry(rehearsalContent)
	if err != nil {
		return "", fmt.Errorf("error parsing rehearsal entry: %w", err)
	}
	if len(newRehearsals) == 0 {
		log.Printf("no rehearsals found %s, %s", conf.Section, rehearsalContent)
	}
	entries, err := stringRehearsalEntries(conf, newRehearsals)
	if err != nil {
		return "", fmt.Errorf("error exporting rehearsal entries: %w", err)
	}

	// The replacement contains code, so it must be literal to avoid expanding text like $1 as a regex group.
	return rehearsalRegex.ReplaceAllLiteralString(input, "## Rehearsal\n"+entries), nil
}

// ReplaceWithLiveLinks replaces the relative links in the input with links to the repo that work.
func ReplaceWithLiveLinks(input, section string, conf *ParseConf) string {
	repoBase := repoAddress + conf.Version
	if conf.Version == "" {
		repoBase = repoAddress + mainBranch
	}

	output := input
	currentDir := repoBase + "/"
	if section != "" {
		output = strings.ReplaceAll(output, "](../", "]("+repoBase+"/")
		currentDir += section + "/"
	}

	return strings.ReplaceAll(output, "](./", "]("+currentDir)
}
