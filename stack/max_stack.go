package stack

import "errors"

// MaxStack is a stack that can return the maximum of the integers pushed.
type MaxStack struct {
	stack1, stack2 []int
}

// ErrEmptyStack occurs when a trying to pop a stack that is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Max solves the problem in O(1) time and O(n) space.
func (maxStack *MaxStack) Max() int {
	if len(maxStack.stack2) == 0 {
		return -1
	}
	return maxStack.stack2[len(maxStack.stack2)-1]
}

// Push adds an integer to the stack.
func (maxStack *MaxStack) Push(i int) {
	maxStack.stack1 = append(maxStack.stack1, i)
	if len(maxStack.stack2) == 0 {
		maxStack.stack2 = append(maxStack.stack2, i)
	} else {
		maxStack.stack2 = append(maxStack.stack2, max(i, maxStack.stack2[len(maxStack.stack2)-1]))
	}
}

// Pop returns the last inserted integer.
func (maxStack *MaxStack) Pop() (int, error) {
	if len(maxStack.stack1) == 0 {
		return -1, ErrEmptyStack
	}
	tmp := maxStack.stack1[len(maxStack.stack1)-1]
	maxStack.stack1 = maxStack.stack1[0 : len(maxStack.stack1)-1]
	maxStack.stack2 = maxStack.stack2[0 : len(maxStack.stack2)-1]
	return tmp, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
