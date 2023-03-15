package stack

import (
	"testing"
)

func TestMaxStack(t *testing.T) {
	var tests = []struct {
		push []int
		pop  int
		max  int
	}{
		{[]int{}, 0, -1},
		{[]int{1, 2, 3, 4, 5}, 2, 3},
		{[]int{1, 2, 3, 4, 5}, 0, 5},
		{[]int{5, 4, 3, 2, 1}, 2, 5},
		{[]int{1, 5, 3, 4}, 1, 5},
	}

	for i, test := range tests {
		stack := new(MaxStack)
		for _, p := range test.push {
			stack.Push(p)
		}

		for i := 0; i < test.pop; i++ {
			stack.Pop()
		}

		got := stack.Max()
		if got != test.max {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.max, got)
		}
	}
}

func TestEmptyStackError(t *testing.T) {
	stack := new(MaxStack)
	if _, err := stack.Pop(); err != ErrEmptyStack {
		t.Fatalf("Calling Pop on an empty stack did not result in empty stack error. Want %#v got %#v", ErrEmptyStack, err)
	}
}
