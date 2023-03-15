package string

import (
	"strings"
)

const (
	zero     = 0
	ten      = 10
	hundred  = 100
	thousand = 1000
	million  = 1000000
	billion  = 1000000000
)

var digits = map[int]string{
	1:        "One",
	2:        "Two",
	3:        "Three",
	4:        "Four",
	5:        "Five",
	6:        "Six",
	7:        "Seven",
	8:        "Eight",
	9:        "Nine",
	ten:      "Ten",
	11:       "Eleven",
	12:       "Twelve",
	13:       "Thirteen",
	14:       "Fourteen",
	15:       "Fifteen",
	16:       "Sixteen",
	17:       "Seventeen",
	18:       "Eighteen",
	19:       "Nineteen",
	20:       "Twenty",
	30:       "Thirty",
	40:       "Forty",
	50:       "Fifty",
	60:       "Sixty",
	70:       "Seventy",
	80:       "Eighty",
	90:       "Ninety",
	hundred:  "Hundred",
	thousand: "Thousand",
	million:  "Million",
	billion:  "Billion",
}

// NumberInEnglish returns how a given number would be read in English
func NumberInEnglish(num int) string {
	output := ""
	eachOutput := ""

	if num == zero {
		return "Zero"
	}

	num, eachOutput = outputIfLarger(num, billion, digits[billion])
	output += eachOutput
	num, eachOutput = outputIfLarger(num, million, digits[million])
	output += eachOutput
	num, eachOutput = outputIfLarger(num, thousand, digits[thousand])
	output += eachOutput

	eachOutput = threeDigitWord(num)
	output += eachOutput

	output = strings.Replace(output, "  ", " ", -1)
	return strings.TrimSpace(output)
}

func outputIfLarger(num, unit int, word string) (int, string) {
	output := ""
	times := howMany(num, unit)
	if times != -1 {
		output = threeDigitWord(times) + " " + word + " "
		num -= times * unit
	}
	return num, output
}

func threeDigitWord(num int) string {
	output := ""
	hund := howMany(num, hundred)

	if hund != -1 {
		output += digits[hund] + " Hundred "

		num -= hund * hundred
	}

	if v, ok := digits[num]; ok {
		return output + v
	}

	tens := howMany(num, ten)
	if tens != -1 {
		output += digits[tens*ten] + " "
		num -= tens * ten
	}

	output += digits[num]
	return output
}

func howMany(num, level int) int {
	if num < level {
		return -1
	}
	return int(float64(num / level))
}
