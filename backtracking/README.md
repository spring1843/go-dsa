# Backtracking

Backtracking is a commonly used [recursive](../recursion) technique in which the solution is built incrementally while simultaneously checking for problem conditions at each step. If the conditions are unmet, the algorithm backtracks to the previous step and tries an alternative approach. The algorithm terminates with success if the final goal is reached while satisfying the problem conditions at each step. On the other hand, if all alternatives have been exhausted and the algorithm has backtracked to the starting point, it terminates with failure.

Backtracking can be compared to how a person solves a maze or searches for an exit in an unfamiliar mall. This technique guarantees a solution if one exists since it involves trying out every possibility until success or failure is achieved.

## Implementation

Backtracking algorithms are typically implemented in these steps:

1. Prune invalid approaches when possible.
2. Generate a partial solution by iterating through available alternatives.
3. Check the validity of the selected alternative according to the problem conditions and rules.
4. Check for solution completion when required.

After finding a backtracking approach to a problem we can start coding by creating a driver function and a recursive function.

The recursive function should receive as input the current state of the progress and uses it to checks for solution completion and to make the next incremental step by recursively calling itself. The recursive function should also have a way to communicate the solution back to the driver. The driver is responsible for making the first call to the recursive function and interpreting the final solution or failure in finding it back to the caller function. To illustrate this lets solve a problem that we will later solve more efficiently as one of the rehearsals in the [Dynamic Programming](../dp/) section.

Sum Up to a Number: Given a set of integers and a positive integer n return true if there is a subset of numbers that can sum up to n and false otherwise. For example for {1,2,3,4,5} and n=10 it should return true because 2+3+5=10 and for n 50 it should return false because there is no subset that can be summed up to 50.

```Go
package main

func SumUpToNumber(numbers []int, n int) bool {
	return sumpUpToNumberRecursive(numbers, n, 0)
}

func sumpUpToNumberRecursive(numbers []int, n int, index int) bool {
	if n == 0 {
		return true
	}
	if n < 0 || index == len(numbers) {
		return false
	}
	return sumpUpToNumberRecursive(numbers, n-numbers[index], index+1) || // Include the number
		sumpUpToNumberRecursive(numbers, n, index+1) // Exclude the number
}
```

The driver function `SumUpToNumber` prepares the first call to the recursive function by sending it all the input it has received and an index number which starts from 0.

The recursive function `sumpUpToNumberRecursive` then checks to see if we have reached a satisfying problem condition i.e. a final success or failure and returns it. Then it calls itself in two ways by checking to see if the sum can be achieved by including the value at current index or excluding it.

## Complexity

The recursive nature of backtracking has a negative effect on performance. The solution to the Sum Up to a Number in the above section for example has the time complexity of O(2^n).

The time complexity of backtracking algorithms may vary depending on the problem at hand. However, they generally require iterating through possible alternatives and checking for validity at each step. Although backtracking may be the only feasible approach to certain problems, it does not always guarantee an optimal solution. Pruning, which involves eliminating known invalid options before iterating through the alternatives, is an effective technique to improve the time complexity of backtracking algorithms.

## Application

Computers use backtracking in board games to select their next moves. The technique is also applied to graphs and trees through the use of [Depth First Search](../graph/graph#depth-first-search---dfs). It also has applications in object detection in image processing.

## Rehearsal

* [Permutations](./permutations_test.go), [Solution](./permutations.go)
* [Generate Parentheses](./generate_parentheses_test.go), [Solution](./generate_parentheses.go)
* [Phone Letter Combinations](./phone_letter_combinations_test.go), [Solution](./phone_letter_combinations.go)
* [Maze](./maze_test.go), [Solution](./maze.go)
* [Sudoku](./sudoku_test.go), [Solution](./sudoku.go)
* [N Queens](./n_queens_test.go), [Solution](./n_queens.go)
