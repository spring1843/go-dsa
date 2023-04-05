package string

import (
	"strconv"
)

// LookAndTell outputs an array such that each element reads the previous
// element the first element being "1"
func LookAndTell(depth int) []string {
	output := []string{"1"}

	if depth <= 0 {
		return []string{}
	}

	if depth == 1 {
		return output
	}

	for i := 1; i < depth; i++ {
		output = append(output, tell(output[i-1]))
	}
	return output
}

func tell(n string) string {
	var step, what string
	var count, until int
	for until < len(n) {
		count, what, until = findHowMany(n, until)
		step += strconv.Itoa(count) + what
	}
	return step
}

func findHowMany(n string, until int) (int, string, int) {
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
