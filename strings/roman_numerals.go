package strings

type romanNumeral struct {
	value  int
	symbol string
}

var romanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// IntToRoman solves the problem in O(n) time and O(n) space.
func IntToRoman(input int) string {
	roman := ""
	for input > 0 {
		for _, romanNumeral := range romanNumerals {
			if input >= romanNumeral.value {
				roman += romanNumeral.symbol
				input -= romanNumeral.value
				break
			}
		}
	}
	return roman
}
