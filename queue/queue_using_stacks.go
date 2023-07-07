package queue

import "errors"

// ErrStackEmpty occurs when popping an empty stack.
var ErrStackEmpty = errors.New("empty stack")

type (
	// UsingStacks is a queue that is made using two stacks.
	UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	// Stack is a stack of integers.
	Stack struct {
		stack []int
	}
)

// enqueue solves the problem in O(1) time and O(n) space.
func (usingStacks *UsingStacks) enqueue(n int) {
	usingStacks.stack1.push(n)
}

// dequeue solves the problem in O(1) time and O(n) space.
func (usingStacks *UsingStacks) dequeue() int {
	if len(usingStacks.stack2.stack) == 0 {
		for len(usingStacks.stack1.stack) != 0 {
			usingStacks.stack2.push(usingStacks.stack1.pop())
		}
	}
	return usingStacks.stack2.pop()
}

func (stack *Stack) push(element int) {
	stack.stack = append(stack.stack, element)
}

func (stack *Stack) pop() int {
	tmp := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return tmp
}
