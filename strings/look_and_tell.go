package strings

import (
	"strconv"
)

// LookAndTell solves the problem in O(n) time and O(n) space.
func LookAndTell(depth int) []string {
	output := []string{"1"}

	if depth <= 0 {
		return []string{"-1"}
	}

	if depth == 1 {
		return output
	}

	for i := 1; i < depth; i++ {
		output = append(output, generateNext(output[i-1]))
	}
	return output
}

// generateNext generates the next number in the sequence.
func generateNext(n string) string {
	var step, what string
	var count, until int
	for until < len(n) {
		count, what, until = countConsequtiveDigits(n, until)
		step += strconv.Itoa(count) + what
	}
	return step
}

// counts the consecutive occurrences of a digit in a string,from a given position
// returning count, digit and the new position
func countConsequtiveDigits(n string, until int) (int, string, int) {
	last := ""
	count := 0
	for until < len(n) {
		char := n[until : until+1]
		if last == "" {
			last = char
			count++
			until++
			continue
		}
		if char != last {
			return count, last, until
		}
		count++
		until++
	}
	return count, last, until
}
