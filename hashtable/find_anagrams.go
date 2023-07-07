package hashtable

import "sort"

type sortRunes []rune

// FindAnagrams solves the problem in O(n*log(n)) time and O(n) space.
func FindAnagrams(words []string) [][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		sorted := sortLetters(word)
		anagrams[sorted] = append(anagrams[sorted], word)
	}
	return toStrings(anagrams)
}

func sortLetters(word string) string {
	r := []rune(word)
	sort.Sort(sortRunes(r))
	return string(r)
}

func toStrings(anagrams map[string][]string) [][]string {
	output := [][]string{}
	for _, anagram := range anagrams {
		if len(anagram) > 1 {
			output = append(output, anagram)
		}
	}
	return output
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
