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

// hash turns a string into a number representing a bitmask of characters
// the output for "abc", "acb", "cba", "aabc", "aaabbcc" and etc… are all the same.
func hash(input string) int64 {
	var res int64
	for _, char := range input {
		res |= 1 << (char - 'a')
	}
	return res
}
