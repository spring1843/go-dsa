# Stack

Stacks are data structures that operate on the principle of Last-In-First-Out (LIFO), where the last element added to the stack is the first to be removed. Stacks allow push, pop and top operations to add, and remove data.

One way to visualize stacks is to think of a backpack where items are placed and later removed in reverse order, with the last item added to the bag being the first item removed.

## Implementation

In Go, stacks can be implemented using slices, which are a dynamic data structure that can easily perform push and pop operations. Alternatively, linked lists can also be used to implement stacks, which can make addition and deletion less expensive.

To push an item onto a stack, the built-in append function can be used, while popping an item can be done by resizing the slice using selectors. An example below demonstrates these operations, showing how a stack of integers `{1, 2, 3}` can be filled and emptied.

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
		return -1, errors.New("Stack is empty")
	}
	tmp := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return tmp, nil
}
```

## Complexity

The push and pop operations in stacks are considered O(1) operations, making them highly efficient. Additionally, many machines have built-in stack instruction sets, further increasing their speed and performance. This unique efficiency and usefulness of stacks have solidified their place as one of the most fundamental data structures, second only to [arrays](../array).

## Application

Stacks are helpful when LIFO operations are desired. Many [graph](../graph) problems are solved with stacks.

Within the context of process execution, a portion of memory known as the "stack" is reserved to hold stack frames. Whenever a function is called, relevant data such as parameters, local variables, and return values are stored within a frame to be accessed after the function has completed. As such, when an excessive number of function calls or an infinite recursive function are made, the computer's ability to store all of this information is exceeded, resulting in the well known stack overflow error.

## Rehearsal

### Max Stack

Implement a stack that can return the max of element it contains. [Solution](max_stack.go) [Test](max_stack_test.go)

### Balancing Symbols

Given a set of symbols including `[]{}()`, determine if the input is is balanced or not. [Solution](balancing_symbols.go) [Test](balancing_symbols_test.go)

### Infix to Postfix Conversion

Given an infix expression e.g. `1*2+3+4*5`, convert it to a postfix expression like `1 2 * 3 + 4 5 *` supporting the four basic arithmetic operations and parenthesis. [Solution](infix_to_postfix.go) [Test](infix_to_postfix_test.go)

### Evaluate Postfix

Given a postfix expression like `1 2 3 + *`, calculate the expression e.g. `5`. [Solution](evaluate_postfix.go) [Test](evaluate_postfix_test.go)

### Basic Calculator

Given an expression containing integers, parentheses and the four basic arithmetic operations like `1*2+3+4*5` calculate the expression into a numerical value like 25. [Solution](basic_calculator.go) [Test](basic_calculator_test.go)
