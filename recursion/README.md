# Recursion

Recursion is a computational technique that implements a [divide-and-conquer](../dnc) approach to problem-solving by breaking down a complex problem into smaller sub-problems. It consists of two components:

* One or more base cases that provide output for simple inputs and terminate recursion
* A recursive case that combines the outputs obtained from recursive function calls to generate a solution for the original problem.

Although recursions enhance code readability, they are usually inefficient and challenging to debug. Consequently, unless they provide a more efficient solution to a problem, such as in the case of [quicksort](../dnc/quick_sort_test.go), they are generally not preferred.

During execution, a program typically stores function variables in a memory area known as the stack before executing recursion. The recursive function may assign different values to the same variables during each recursion. When the recursion ends, the stack pops and remembers the values. However, the stack will grow with each call if recursion continues indefinitely, causing the familiar stack overflow error. Since recursion employs the stack to execute, every recursive problem can be converted into an iterative one. This transformation, however, typically leads to more complex code and may require a [stack](../stack).

## Implementation

The computation of the nth Fibonacci number can be achieved with recursion. For the base cases, it is established that the Fibonacci number for 0 and 1 is respectively 0 and 1. Additionally, for any other number, the Fibonacci number equals the sum of the two preceding Fibonacci numbers.

```Go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n-2)
}
```

When formulating recursive algorithms, it is essential to consider the following four rules of recursion:

1. It is imperative to establish a base case, or else the program will terminate abruptly
2. The algorithm should progress toward the base case at each recursive call.
3. Recursive calls are presumed effective; thus, traversing every recursive call and performing bookkeeping is unnecessary.
4. Use memoization, a technique that prevents redundant computation by caching previously computed results, can enhance the algorithm's efficiency.

## Complexity

Recursions are often inefficient in both time and space complexity. The number of recursive calls required to solve a problem can grow exponentially (for example the Fibonacci implementation in the last section). In worst cases, such as [permutations](../backtracking/ permutations_test.go) they can also be O(n!).

There are a few different ways of determining the time complexity of recursive algorithms:

1. Recurrence Relations: This approach involves defining a recurrence relation that expresses the algorithm's time complexity in terms of its sub-problems' time complexity. For example, for the recursive Fibonacci algorithm, the recurrence relation is T(n) = T(n-1) + T(n-2) + O(1), where T(n) represents the time complexity of the algorithm for an input of size n.
2. Recursion Trees: This method involves drawing a tree to represent the algorithm's recursive calls. The algorithm's time complexity can be calculated by summing the work done at each level of the tree. For example, for the recursive factorial algorithm, each level of the tree represents a call to the function with a smaller input size, and the work done at each level is constant.
3. Master Theorem: This approach is a formula for solving recurrence relations that have the form T(n) = aT(n/b) + f(n). The Master Theorem can be used to quickly determine the time complexity of some [Divide-and-conquer](../dnd) algorithms.

The space complexity of recursive calls is affected by having to store a copy of the state and variables in the stack with each recursion.

## Application

Recursion finds practical application within a range of algorithms, including [Dynamic Programming](../dp), [Graph](../graph), [Divide-and-conquer](../dnd), and [Backtracking](../backtracking). Typically, recursion is best suited to problems that exhibit sub-problems or require calculating the nth or first nth value in a series.

## Rehearsal

* [Reverse an integer recursively](./reverse_number_test.go), [Solution](./reverse_number.go)
* [Palindrome](./is_palindrome_test.go), [Solution](./is_palindrome.go)
* [Climbing Stairs](./climbing_stairs_test.go), [Solution](./climbing_stairs.go)
* [Exponentiation](./exponentiation_test.go), [Solution](./exponentiation.go)
* [Multiplication](./multiplication_test.go), [Solution](./multiplication.go)
* [Regular Expressions Matching](./regular_expressions_test.go), [Solution](./regular_expressions.go)
* [Expression Operators](./expression_operators_test.go), [Solution](./expression_operators.go)
