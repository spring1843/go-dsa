package strings

import "testing"

/*
TestIntToRoman tests solution(s) with the following signature and problem description:

	func IntToRoman(input int) string

Given a positive integer like return the equivalent inRoman numerals:

For example:

* Given 1, return I
* Given 4, return IV
* Given 5, return V
* Given 9, return IX
* Given 10, return X
* Given 40, return XL
* Given 50, return L
* Given 90, return XC
* Given 100, return C
* Given 400, return CD
* Given 500, return D
* Given 900, return CM
* Given 1000, return M
* Given 1917, return MCMXVII.
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
