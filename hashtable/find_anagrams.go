package hashtable

import "slices"

// FindAnagrams solves the problem in O(n*Log n) time and O(n) space.
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
	slices.Sort(r)
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
