# Dynamic Programming

Dynamic Programming (DP) and [divide-and-conquer (DNC)](../dnc) share a common strategy of breaking down a problem into smaller sub-problems. However, DP works by solving each sub-problem only once and preemptively eliminating sub-problems that cannot contribute to optimal solutions. This approach makes DP more efficient than DNC where same sub-problems may be solved multiple times.

## Implementation

Dynamic Programming (DP) approach to algorithms typically involves four steps:

1. Characterize an optimal solution.
2. Define the value of an optimal solution through [recursion](../recursion).
3. Compute the value of the optimal solution in a bottom-up manner.
4. Determine the optimal solution using the values computed in the previous steps.

If only the solution value is needed, step 4 may be omitted. Conversely, when the solution itself only is required, it is advisable to consider the necessary values at the final step. This will facilitate storing pertinent information at step 3 and simplify step 4.

There are two general approaches for writing DP algorithms: top-down and bottom-up.

The top-down approach starts with the final solution and breaks it down into smaller sub-problems. The final solution is found by calculating smaller sub-problems. A caching mechanism stores each sub-problem solution and prevents unnecessary re-calculations. This technique is also known as memoization.

The performance of Fibonacci solution we introduced in [recursion](../recursion/README.md) can be significantly improved by memoization.

```Go
package main

import (
	"fmt"
)

var fib = make(map[int]int)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if result, ok := fib[n]; ok {
		return result
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```

The bottom-up approach builds the solution iteratively from smaller sub-problems towards the final solution.

```Go
package main

import (
	"fmt"
)

func fibonacci(n int) int {
	var fib = map[int]int{
		0: 0,
		1: 1,
	}

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
```

## Complexity

The complexity of DP algorithms depends on the size and structure of the sub-problems. The solution "House Robber" problem in rehearsal section is an example of a DP bottom-up solution with O(n) complexity.

A top-down and bottom-up solutions to the same problem usually have the same time complexity. For example both the top-down and bottom-up solutions to the Fibonacci problem have the time complexity of O(n).

## Application

DP is well-suited for tackling complex problems, including logistics, game theory, machine learning, science, resource allocation, and investment policies. In [graph](../graph/) theory, DP is commonly used to determine the shortest path between two points. DP algorithms are particularly effective at optimizing problems with many decisions that can be made throughout multiple stages, where the goal is to identify the most optimal set of decisions according to a criteria.

## Rehearsal

* [Rod Cutting](./rod_cutting_test.go), [Solution](./rod_cutting.go)
* [Sum Up to Number](./sum_up_to_integer_test.go), [Solution](./sum_up_to_integer.go)
* [House Robber](./house_robber_test.go), [Solution](./house_robber.go).
* [Interleaving String](./interleaving_string_test.go), [Solution](./interleaving_string.go)
* [Minimum Deletion to Make a Palindrome](./minimum_deletion_to_make_palindrome_test.go), [Solution](./minimum_deletion_to_make_palindrome.go)
* [Word Distance](./word_distance_test.go), [Solution](./word_distance.go)
