package problems

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewRehearsalEntry(t *testing.T) {
	input := `## Rehearsal

* [Bubble Sort](./bubble_sort_test.go), [Solution](./bubble_sort.go)
* [Reverse Array In-place](./reverse_inplace_test.go), [Solution](./reverse_inplace.go)
* [Add Two Numbers](./add_two_numbers_test.go), [Solution](./add_two_numbers.go)
* [Find Duplicate in Array](./find_duplicate_in_array_test.go), [Solution](./find_duplicate_in_array.go)
* [Zero Sum Triplets](./zero_sum_triplets_test.go), [Solution](./zero_sum_triplets.go)
* [Product of All Other Elements](./product_of_all_other_elements_test.go), [Solution](./product_of_all_other_elements.go)`

	entries, err := newRehearsalEntry(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 6 {
		t.Fatalf("expected 5 entries, got %d", len(entries))
	}
}

func TestParseSection(t *testing.T) {
	const (
		readme   = "# Example\n\n## Rehearsal\n\n- [Example](./example_test.go), [Solution](./example.go)\n"
		testFile = "package example\n\n// A 1 inch rod is $1.\n"
		solution = "package example\n\nfunc Example() {}\n"
	)

	dir := t.TempDir()
	section := filepath.Join(dir, "example")
	if err := os.Mkdir(section, 0o750); err != nil {
		t.Fatal(err)
	}
	for name, content := range map[string]string{
		"README.md":       readme,
		"example_test.go": testFile,
		"example.go":      solution,
	} {
		if err := os.WriteFile(filepath.Join(section, name), []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}

	conf := newParseConf(dir, "example", "")
	got, err := ParseSection(conf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, want := range []string{testFile, solution} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected output to contain %q, got %q", want, got)
		}
	}

	for _, name := range []string{"example.go", "example_test.go"} {
		if err := os.Remove(filepath.Join(section, name)); err != nil {
			t.Fatal(err)
		}
		if _, err := ParseSection(conf); !errors.Is(err, fs.ErrNotExist) {
			t.Fatalf("expected a missing file error after removing %s, got %v", name, err)
		}
	}
}

func TestReplaceWithLiveLinks(t *testing.T) {
	tests := []struct {
		input, section, version, want string
	}{
		{"[a](./preface.md)", "", "", "[a](" + repoAddress + "main/preface.md)"},
		{"[a](./heap_sort.go)", "heap", "", "[a](" + repoAddress + "main/heap/heap_sort.go)"},
		{"[a](../tree)", "heap", "v1", "[a](" + repoAddress + "v1/tree)"},
	}

	for i, test := range tests {
		if got := ReplaceWithLiveLinks(test.input, test.section, newParseConf("", "", test.version)); got != test.want {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.want, got)
		}
	}
}

func newParseConf(dir, section, version string) *ParseConf {
	return &ParseConf{
		Dir:                    dir,
		Section:                section,
		ReplacePackageWithMain: false,
		ReplaceWithLiveLinks:   false,
		Version:                version,
	}
}
