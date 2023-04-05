package string

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
	ErrPopStack = errors.New("can not Pop on an empty stack")
)

// ReverseVowels reverses the order of vowels in a string
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
			newChar, err := pop()
			if err != nil {
				return "", err
			}
			reversed += newChar
			continue
		}
		reversed += string(r)
	}
	return reversed, nil
}

func push(v string) {
	stack = append(stack, v)
}

func pop() (string, error) {
	if len(stack) == 0 {
		return "", ErrPopStack
	}
	tmp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return tmp, nil
}
