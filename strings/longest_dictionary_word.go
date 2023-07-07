package strings

// LongestDictionaryWordContainingKey solves the problem in O(n) time and O(1) space.
func LongestDictionaryWordContainingKey(key string, dic []string) string {
	keyNum := hash(key)
	longest := ""

	for _, word := range dic {
		wordNum := hash(word)
		if keyNum&wordNum == keyNum {
			if len(longest) < len(word) {
				longest = word
			}
		}
	}
	return longest
}

// hash turns a string into a number
// the output for "abc", "acb", "cba" and etc... are all the same.
func hash(s string) rune {
	var res rune
	for _, w := range s {
		res |= w
	}
	return res
}
