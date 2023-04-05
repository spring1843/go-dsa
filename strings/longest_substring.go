package strings

// LongestSubstringOfTwoChars returns the longest substring of two different characters in a string
func LongestSubstringOfTwoChars(input string) string {
	if len(input) <= 1 {
		return ""
	}

	best, pos := 0, 0
	ch1, ch2 := input[0:1], input[1:1]

	start := 0
	for end := 1; end < len(input); end++ {
		if string(input[end]) != ch1 && string(input[end]) != ch2 {
			for start = end - 1; start > 0; start-- {
				if string(input[start-1]) != string(input[end-1]) {
					break
				}
			}
			ch1 = string(input[start])
			ch2 = string(input[end])
		}

		// More than two unique characters are found
		if best < end-start+1 {
			best = end - start + 1
			pos = start
		}
	}

	return input[pos : pos+best]
}
