package problems

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func ParseSection(dir, section string) (string, error) {
	content, err := os.ReadFile(filepath.Join(dir, section, "README.md"))
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	preparedContent, err := replaceRehearsal(string(content), dir, section)
	if err != nil {
		return "", fmt.Errorf("error replacing the rehearsal section in file: %w", err)
	}

	return preparedContent, nil
}

func replaceRehearsal(input, dir, section string) (string, error) {
	rehearsalRegex := regexp.MustCompile(`(?s)(## Rehearsal\n.*?)(\n##|\z)`)

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
		log.Printf("no rehearsals found %s, %s", section, rehearsalContent)
	}
	replacement := fmt.Sprintf("## Rehearsal\n%s", stringRehearsalEntries(dir, section, newRehearsals))
	return rehearsalRegex.ReplaceAllString(input, replacement), nil
}
