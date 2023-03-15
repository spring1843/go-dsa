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

// PhoneLetterCombinations intakes the digits from 2 to 9 that represent phone buttons
// and returns all possible combinations of letters that could be generated from those.
func PhoneLetterCombinations(digits string) []string {
	combinations := make([]string, 0)
	if digits == "" {
		return combinations
	}
	phoneLetterCombinationsRecursive(digits, 0, &combinations, "")
	return combinations
}

func phoneLetterCombinationsRecursive(digits string, p int, combinations *[]string, s string) {
	if p == len(digits) {
		*combinations = append(*combinations, s)
		return
	}

	for _, digit := range phone[digits[p]] {
		phoneLetterCombinationsRecursive(digits, p+1, combinations, s+string(digit))
	}
}
