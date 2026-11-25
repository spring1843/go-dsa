package strings

import (
	"strings"
)

var (
	stack  []string
	vowels = map[string]bool{
		`a`: true,
		`e`: true,
		`o`: true,
		`i`: true,
		`u`: true,
	}
)

// ReverseVowels solves the problem in O(n) time and O(n) space.
func ReverseVowels(str string) string {
	stack = []string{}
	var reversed strings.Builder
	for _, r := range str {
		if _, ok := vowels[string(r)]; ok {
			push(string(r))
		}
	}
	for _, r := range str {
		if _, ok := vowels[string(r)]; ok {
			reversed.WriteString(pop())
			continue
		}
		reversed.WriteRune(r)
	}
	return reversed.String()
}

func push(v string) {
	stack = append(stack, v)
}

func pop() string {
	tmp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return tmp
}
