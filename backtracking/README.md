# Backtracking

Backtracking is a commonly used [recursive](../recursion) technique in which the solution is built incrementally while simultaneously checking for problem conditions at each step. If the conditions are not met, the algorithm backtracks to the previous step and tries an alternative approach. The algorithm terminates with success if the final goal is reached while satisfying the problem conditions at each step. On the other hand, if all alternatives have been exhausted and the algorithm has backtracked to the starting point, it terminates with failure.

Backtracking can be compared to how a person solves a maze or searches for an exit in an unfamiliar mall. This technique guarantees a solution if one exists since it involves trying out every possibility until success or failure is achieved.

## Implementation

Backtracking algorithms are typically implemented in these steps:

1. Prune invalid approaches when possible.
2. Generate a partial solution by iterating through available alternatives.
3. Check the validity of the selected alternative according to the problem conditions and rules.
4. Check for solution completion when required.

## Complexity

The time complexity of backtracking algorithms may vary depending on the problem at hand, however they generally require iterating through possible alternatives and checking for validity at each step. Although backtracking may be the only feasible approach to certain problems, it does not always guarantee an optimal solution. To improve the time complexity of backtracking algorithms, pruning, which involves eliminating known invalid options before iterating through the alternatives, is an effective technique.

## Application

Backtracking is widely used to solve board games and computers use it to select their next moves. Furthermore, the backtracking technique is also applied to graphs and trees through the use of [Depth First Search](../graph/graph#depth-first-search---dfs). It also has applications in object detection and image processing.

## Rehearsal

## Permutations

* [Permutations](permutations_test.go),  [Solution](permutations.go)
* [Generate Parentheses](./generate_parentheses_test.go), [Solution](generate_parentheses.go)
* [Phone Letter Combinations](phone_letter_combinations_test.go), [Solution](phone_letter_combinations.go)
* [Sudoku](sudoku_test.go), [Solution](sudoku.go)
* [N Queens](n_queens_test.go), [Solution](n_queens.go)
