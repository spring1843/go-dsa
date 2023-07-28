# Greedy Algorithms

The concept of greedy algorithms involves choosing a naive solution and continuously refining it rather than developing a sophisticated algorithm by beginning with what appears to be an easy win. This approach chooses a local optimum, aiming to build upon it to ultimately arrive at the global optimum.

For example, consider someone kayaking to an island and discovering a treasure while a tsunami is imminent, and they want to take some pieces of treasure back with them. Given their limit in time and transportation, which items should they take to maximize the value? The greedy approach would involve selecting the most ostensibly valuable items first, such as heavy pieces of gold or diamonds rather than small pieces of silver. However, if a valuable item, such as a small silver ring with a significant archeological value that makes it the most precious piece, is overlooked, this approach may result in a suboptimal solution. This is similar to the [knapsack](./knapsack_test.go) problem.

Therefore, greed and greedy algorithms may not always produce optimal solutions if the local optimum does not equal the global optimum. However, they can be useful for approximating solutions where exact answers are not required.

## Implementation

The process for developing a greedy algorithm can be broken down into six steps:

1. Identify the optimal problem structure.
2. Develop a recursive solution.
3. Show that with a greedy choice, only one subproblem remains.
4. Prove that it is safe to make the greedy choice.
5. Write a recursive algorithm implementing the greedy solution.
6. Convert the recursive algorithm into an iterative one.

## Complexity

The time complexity of a greedy algorithm depends on the specific problem being solved and the complexity of each step in the algorithm.

## Application

Greedy algorithms can be applied to optimization problems in which the aim is to optimize a specific objective function or maximize or minimize a certain value.

## Rehearsal

* [Maximum Stock Profit](./max_stock_profit_test.go), [Solution](./max_stock_profit.go)
* [Activity Selector](./activity_selector_test.go), [Solution](./activity_selector.go)
* [Knapsack](./knapsack_test.go), [Solution](./knapsack.go)
* [Jump Game](./jump_game_test.go), [Solution](./jump_game.go)
* [Maximum Number](./max_number_test.go), [Solution](./max_number.go)
* [Task Scheduling](./task_scheduling_test.go), [Solution](./task_scheduling.go)
