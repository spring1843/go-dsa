package strings

import "strings"

// LongestDictionaryWordContainingKey solves the problem in O(n) time and O(1) space.
func LongestDictionaryWordContainingKey(key string, dic []string) string {
	keyCharCount := make(map[rune]int)
	for _, char := range key {
		keyCharCount[char]++
	}

	var longest string
	for _, word := range dic {
		if containsAllChars(word, keyCharCount) && len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// function to check if word contains all characters in the key 
func containsAllChars(word string, keyCharCount map[rune]int) bool {
	for char, count := range keyCharCount {
		if strings.Count(word, string(char)) < count {
			return false
		}
	}
	return true
}
