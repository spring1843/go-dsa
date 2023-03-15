package queue

import "errors"

// ErrStackEmpty occurs when popping an empty stack
var ErrStackEmpty = errors.New("Empty stack")

type (
	// UsingStacks is a queue that is made using two stacks
	UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	// Stack is a stack of integers
	Stack struct {
		stack []int
	}
)

func (usingStacks *UsingStacks) enqueue(n int) {
	usingStacks.stack1.push(n)
}

func (usingStacks *UsingStacks) dequeue() (int, error) {
	if len(usingStacks.stack2.stack) == 0 {
		for len(usingStacks.stack1.stack) != 0 {
			n, err := usingStacks.stack1.pop()
			if err != nil {
				return 0, err
			}
			usingStacks.stack2.push(n)
		}
	}
	return usingStacks.stack2.pop()
}

func (stack *Stack) push(element int) {
	stack.stack = append(stack.stack, element)
}

func (stack *Stack) pop() (int, error) {
	if len(stack.stack) == 0 {
		return 0, ErrStackEmpty
	}
	tmp := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return tmp, nil
}
