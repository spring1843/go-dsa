# Stack

Stacks are data structures that operate on the Last-In-First-Out (LIFO) principle, where the last element added to the stack is the first to be removed. Stacks allow push, pop, and top operations to add, remove, and read data.

One way to visualize stacks is to think of a backpack where items are placed and later removed in reverse order, with the last item added to the bag being the first item removed.

## Implementation

In Go, stacks can be implemented using doubly [linked lists](../linkedlist/) or [arrays and slices](../array/). Here is a linked list implementation:

```Go
package main

import (
	"container/list"
	"errors"
)

var stack = list.New()

func main() {
	push(1) // [1]
	push(2) // [1,2]
	push(3) // [1,2,3]
	pop()   // [1,2]
	pop()   // [1]
	pop()   // []
}

func push(val int) {
	stack.PushBack(val)
}

func pop() (int, error) {
	if stack.Len() == 0 {
		return -1, errors.New("stack is empty")
	}
	return stack.Remove(stack.Back()).(int), nil
}
```

In the slice implementation to push an item onto a stack, the built-in append function can be used, while popping an item can be done by resizing the slice using selectors. An example below demonstrates these operations, showing how a stack of integers `{1, 2, 3}` can be filled and emptied.

```Go
package main

import "errors"

var stack []int

func main() {
	push(1) // [1]
	push(2) // [1,2]
	push(3) // [1,2,3]
	pop()   // [1,2]
	pop()   // [1]
	pop()   // []
}

func push(val int) {
	stack = append(stack, val)
}

func pop() (int, error) {
	if len(stack) == 0 {
		return -1, errors.New("stack is empty")
	}
	tmp := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return tmp, nil
}
```

## Complexity

Push and pop operations in stacks are considered O(1) operations, making them highly efficient. Additionally, many machines have built-in stack instruction sets, further increasing their performance. Stacks' unique efficiency and usefulness have solidified their place as one of the most fundamental data structures, second only to [arrays](../array).

## Application

Stacks are helpful when LIFO operations are desired. Many [graph](../graph) problems are solved with stacks.

During process execution, a portion of memory known as the "stack" is reserved to hold stack frames. Whenever a function is called, relevant data such as parameters, local variables, and return values are stored within a frame to be accessed after the function has been completed. When an excessive number of function calls or an infinite recursive function are made, the computer's ability to store all of this information is exceeded. This results in the well-known stack overflow error.

## Rehearsal

* [Max Stack](./max_stack_test.go), [Solution](./max_stack.go)
* [Balancing Symbols](./balancing_symbols_test.go), [Solution](./balancing_symbols.go)
* [Infix to Postfix Conversion](./infix_to_postfix_test.go), [Solution](./infix_to_postfix.go)
* [Evaluate Postfix](./evaluate_postfix_test.go), [Solution](./evaluate_postfix.go)
* [Longest Valid Parentheses](./longest_valid_parentheses_test.go), [Solution](./longest_valid_parentheses.go)
* [Basic Calculator](./basic_calculator_test.go), [Solution](basic_calculator.go)
