package backtracking

var phone = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// PhoneLetterCombinations solves the problem in O(3^n) time and O(3^n) space.
func PhoneLetterCombinations(digits string) []string {
	combinations := []string{}
	if len(digits) > 0 {
		phoneLetterCombinationsRecursive(digits, "", &combinations)
	}
	return combinations
}

func phoneLetterCombinationsRecursive(digits, prefix string, combinations *[]string) {
	if len(prefix) == len(digits) {
		*combinations = append(*combinations, prefix)
		return
	}

	for _, letter := range phone[digits[len(prefix)]] {
		phoneLetterCombinationsRecursive(digits, prefix+string(letter), combinations)
	}
}
