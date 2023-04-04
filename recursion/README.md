# Recursion

Recursion is a computational technique that implements a [divide and conquer](../dnc) approach to problem-solving by breaking down a complex problem into smaller sub-problems. It consists of two components:

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

1. Recurrence Relations: This approach involves defining a recurrence relation that expresses the time complexity of the algorithm in terms of the time complexity of its sub-problems. For example, for the recursive Fibonacci algorithm, the recurrence relation is T(n) = T(n-1) + T(n-2) + O(1), where T(n) represents the time complexity of the algorithm for an input of size n.
2. Recursion Trees: This method involves drawing a tree to represent the recursive calls made by the algorithm. The time complexity of the algorithm can be calculated by summing the work done at each level of the tree. For example, for the recursive factorial algorithm, each level of the tree represents a call to the function with a smaller input size, and the work done at each level is constant.
3. Master Theorem: This approach is a formula for solving recurrence relations that have the form T(n) = aT(n/b) + f(n). The Master Theorem can be used to quickly determine the time complexity of some [Divide-and-conquer](../dnd) algorithms.

## Application

Recursion finds practical application within a range of algorithms, including [Dynamic Programming](../dp), [Graph](../graph), [Divide-and-conquer](../dnd), and [Backtracking](../backtracking). Typically, recursion is best suited to problems that exhibit sub-problems or require the calculation of the nth or first nth value in a series.

## Rehearsal

### Reverse an integer recursively

Given an integer like 321 return a reversed number using recursion where the same digits are repeated in the reverse order like 321. [Solution](reverse_number.go) [Test](reverse_number_test.go)

### Palindrome

Given a string like `abba` return true if it's a palindrome and false otherwise. [Solution](is_palindrome.go) [Test](is_palindrome_test.go)

### Climbing Stairs

Given n the number of steps, return in how many ways you can climb these stairs if you are only able to climb 1 or 2 steps at a time. [Solution](climbing_stairs.go) [Test](climbing_stairs_test.go)

### Exponentiation

Given x and n, return x raised to the power of n in an efficient manner. [Solution](exponentiation.go) [Test](exponentiation_test.go)

### Permutations

Given a set of integers like `{1,2}`, return all possible permutations like `{1,2},{2,1}`.  [Solution](permutations.go) [Test](permutations_test.go)

### Regular Expressions Matching

Given an input and a regular expression pattern where `.` denotes to any character and `*` denotes to zero or more of the proceeding characters, write a recursive function to return true if the input matches the pattern and false otherwise. [Solution](regular_expressions.go) [Test](regular_expressions_test.go)
