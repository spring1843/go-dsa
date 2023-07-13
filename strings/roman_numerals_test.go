package strings

import "testing"

/*
TestIntToRoman tests solution(s) with the following signature and problem description:

	func IntToRoman(input int) string

Given a positive integer like 1917 return the Roman numeral like MCMXVII.
*/
func TestIntToRoman(t *testing.T) {
	tests := []struct {
		number int
		roman  string
	}{
		{0, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{58, "LVIII"},
		{60, "LX"},
		{1000, "M"},
		{1917, "MCMXVII"},
	}
	for i, test := range tests {
		if got := IntToRoman(test.number); got != test.roman {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.roman, got)
		}
	}
}
