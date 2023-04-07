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

func (usingStacks *UsingStacks) enqueue(n int) {
	usingStacks.stack1.push(n)
}

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
	if len(stack.stack) == 0 {
		return 0
	}
	tmp := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return tmp
}
