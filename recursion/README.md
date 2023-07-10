# Recursion

Recursion is a computational technique that implements a [divide and conquer](../dnc) approach to problem-solving by breaking down a complex problem into smaller subproblems. It consists of two components:

* one or more base cases that provide output for simple inputs
* a recursive case that combines the outputs obtained from recursive function calls to generate a solution for the original problem.

Although recursions enhance code readability, they are usually not efficient and can be challenging to debug. Consequently, unless they provide a more efficient solution to a problem, such as in the case of Quicksort, they are generally not preferred.

During execution, a program typically stores function variables in a memory area known as the stack before executing recursion. The recursive function may assign different values to the same variables during each recursion. When the recursion ends, the stack pops and remembers the values. However, if recursion continues indefinitely, the stack will grow with each call causing the familiar stack overflow error. Since recursion employs the stack to execute, every recursive problem can be converted into an iterative one. This transformation, however, typically leads to more complex code and may require the use of a [stack](../stack).

## Implementation

The computation of the nth Fibonacci number can be achieved recursively. For the base cases, it is established that the Fibonacci number for 0 and 1 is 0 and 1, respectively. Additionally, for any other number, the Fibonacci number is equal to the sum of the two preceding Fibonacci numbers.

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
2. The algorithm should make progress towards the base case at each recursive call
3. Recursive calls are presumed to be effective; thus, it is unnecessary to traverse every recursive call and perform bookkeeping
4. Use memoization, a technique that prevents redundant computation by caching previously computed results, can be used to enhance the algorithm's efficiency

## Complexity

Recursions are often inefficient, both in terms of time and memory usage. The number of recursive calls required to solve a problem can grow exponentially, reaching a complexity of n factorial in some cases. For instance, the recursive Fibonacci algorithm has a time complexity of O(2^n).

There are a few different ways of determining the time complexity of recursive algorithms:

1. Recurrence Relations: This approach involves defining a recurrence relation that expresses the time complexity of the algorithm in terms of the time complexity of its subproblems. For example, for the recursive Fibonacci algorithm, the recurrence relation is T(n) = T(n-1) + T(n-2) + O(1), where T(n) represents the time complexity of the algorithm for an input of size n.
2. Recursion Trees: This method involves drawing a tree to represent the recursive calls made by the algorithm. The time complexity of the algorithm can be calculated by summing the work done at each level of the tree. For example, for the recursive factorial algorithm, each level of the tree represents a call to the function with a smaller input size, and the work done at each level is constant.
3. Master Theorem: This approach is a formula for solving recurrence relations that have the form T(n) = aT(n/b) + f(n). The Master Theorem can be used to quickly determine the time complexity of some [Divide-and-conquer](../dnd) algorithms.

## Application

Recursion finds practical application within a range of algorithms, including [Dynamic Programming](../dp), [Graph](../graph), [Divide-and-conquer](../dnd), and [Backtracking](../backtracking). Typically, recursion is best suited to problems that exhibit subproblems or require the calculation of the nth or first nth value in a series.

## Rehearsal

* [Reverse an integer recursively](reverse_number_test.go), [Solution](reverse_number.go)
* [Palindrome](is_palindrome_test.go), [Solution](is_palindrome.go)
* [Climbing Stairs](climbing_stairs_test.go), [Solution](climbing_stairs.go)
* [Exponentiation](exponentiation_test.go), [Solution](exponentiation.go)
* [Regular Expressions Matching](regular_expressions_test.go), [Solution](regular_expressions.go)
