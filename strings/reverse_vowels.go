package strings

import "errors"

var (
	stack  []string
	vowels = map[string]bool{
		`a`: true,
		`e`: true,
		`o`: true,
		`i`: true,
		`u`: true,
	}
	// ErrPopStack is returned when the stack is empty.
	ErrPopStack = errors.New("can not Pop on an empty stack")
)

// ReverseVowels solves the problem in O(n) time and O(n) space.
func ReverseVowels(str string) (string, error) {
	stack = []string{}
	reversed := ""
	for _, r := range str {
		if _, ok := vowels[string(r)]; ok {
			push(string(r))
		}
	}
	for _, r := range str {
		if _, ok := vowels[string(r)]; ok {
			reversed += pop()
			continue
		}
		reversed += string(r)
	}
	return reversed, nil
}

func push(v string) {
	stack = append(stack, v)
}

func pop() string {
	tmp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return tmp
}
