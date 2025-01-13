# Backtracking

Backtracking is a commonly used [recursive](../recursion) technique in which the solution is built incrementally while simultaneously checking for problem conditions at each step. If the conditions are unmet, the algorithm backtracks to the previous step and tries an alternative approach. The algorithm terminates with success if the final goal is reached while satisfying the problem conditions at each step. On the other hand, if all alternatives have been exhausted and the algorithm has backtracked to the starting point, it terminates with failure.

Backtracking can be compared to how a person solves a maze or searches for an exit in an unfamiliar mall. This technique guarantees a solution if one exists since it involves trying out every possibility until success or failure is achieved.

## Implementation

Backtracking algorithms are typically implemented in these steps:

1. Prune invalid approaches when possible.
2. Generate a partial solution by iterating through available alternatives.
3. Check the validity of the selected alternative according to the problem conditions and rules.
4. Check base cases for solution completion when required.

After finding a backtracking approach to a problem we can start coding by creating a driver function and a recursive function.

The recursive function should receive current state of the progress as input and uses it to checks the base cases for solution completion and to make the next incremental step by recursively calling itself. The recursive function should also have a way to communicate the solution back to the driver. The driver is responsible for making the first call to the recursive function and interpreting the final solution or failure in finding it back to the caller function. Lets review an example to illustrate this process.

```Go
package main

import "fmt"

// Contains returns true if text contains pattern and false otherwise. For example:
// Contains("abracadabra", "cad") returns true
// Contains("abracadabra", "ra") returns true
// Contains("abracadabra", "zebra") returns false
func Contains(text, pattern string) bool {
	return containsRecursive(text, pattern, 0, 0)
}

func containsRecursive(text, pattern string, textIndex, patternIndex int) bool {
	if patternIndex == len(pattern) {
		return true 
	}
	if textIndex == len(text) {
		return false 
	}

	if text[textIndex] == pattern[patternIndex] {
		if containsRecursive(text, pattern, textIndex+1, patternIndex+1) {
			return true
		}
	}

	return containsRecursive(text, pattern, textIndex+1, 0)
}
```

To solve this problem using the backtracking technique:

* Recursively iterate through each character of `text` at `textIndex` and compare it with the first character of the `pattern`:
    * If they match compare every remaining character in text with every remaining character in pattern and if they are all equal then success base case condition is satisfied.
    * If they are not equal then  increment the `textIndex` repeat

The driver function `Contains` prepares the first call to the recursive function by sending it all the inputs it has received and to index numbers as `textIndex` and `patternIndex` as 0.

The recursive function `containsRecursive` then checks for base cases. If `patternIndex` equals the length of `pattern` it means every character of the pattern has been compared to a substring of text and they have all matched so the function returns true. If `textIndex`equals the length of `text` then all characters have been explored and the pattern has not been found so it returns false.

Then it checks to see if the character at `textIndex` in `text` matches the character at `patternIndex` in `pattern`:
* If they match it will recursively call the same function to check the remaining characters.
* If they do not match it backtracks to the next character in text by recursively calling the same function with an incremented textIndex.

For example for inputs `abracadabra` and `cad` it will check to see if `a` equals `c`, if it doesn't then it looks at the next character in text which is `b`. Since it does not match it will keep incrementing `textIndex` until it equals 4 at which point it will match `c`. Then it will check to see if text and pattern match at index 5 and 1. If they do then 6 and 2, and if they do since `cad` has three letters and we have found 3 matching elements we can conclude the text contains the pattern.

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
