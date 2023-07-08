# Greedy Algorithms

The concept of greedy algorithms involves choosing a naive solution and continuously refining it, rather than developing a sophisticated algorithm, by beginning with what appears to be an easy win. In this approach, a local optimum is chosen, and the aim is to build upon it to ultimately arrive at the global optimum.

For example, consider someone kayaking to an island and discovering a treasure while a tsunami is imminent. To decide which pieces of treasure to take back, the greedy approach would involve selecting the most seemingly valuable items first, such as heavy pieces of gold or diamonds rather than small pieces of silver. However, this approach may result in a suboptimal solution if a valuable item, such as a small silver ring with significant archeological value that makes it the most valuable piece, is overlooked. This is similar to the knapsack problem solved in the rehearsals.

Therefore, greed and greedy algorithms may not always produce optimal solutions if the local optimum does not equal the global optimum. However, they can be useful for approximating solutions in cases where exact answers are not required.

## Implementation

The process for developing a greedy algorithm can be structured into six steps.

The process for developing a greedy algorithm can be broken down into six steps:

1. Identify the optimal structure of the problem.
2. Develop a recursive solution.
3. Show that with a greedy choice only one subproblem remains.
4. Prove that it is safe to make the greedy choice.
5. Write a recursive algorithm implementing the greedy solution.
6. Convert the recursive algorithm into an iterative one.

## Complexity

The time complexity of a greedy algorithm depends on the specific problem being solved, and the complexity of each step in the algorithm.

## Application

Greedy algorithms can be applied to optimization problems where the aim is to optimize a specific objective function or maximizing or minimizing a certain value.

## Rehearsal

* [Maximum Stock Profit](max_stock_profit_test.go), [Solution](max_stock_profit.go)
* [Activity Selector](activity_selector_test.go), [Solution](activity_selector.go)
* [Knapsack](knapsack_test.go), [Solution](knapsack.go)
* [Jump Game](jump_game_test.go), [Solution](jump_game.go)
* [Task Scheduling](task_scheduling_test.go), [Solution](task_scheduling.go)
