# Greedy Algorithms

The concept of greedy algorithms involves choosing a naive solution and continuously refining it, rather than developing a sophisticated algorithm, by beginning with what appears to be an easy win. In this approach, a local optimum is chosen, and the aim is to build upon it to ultimately arrive at the global optimum.

For example, consider someone kayaking to an island and discovering a treasure while a tsunami is imminent. To decide which pieces of treasure to take back, the greedy approach would involve selecting the most seemingly valuable items first, such as heavy pieces of gold or diamonds rather than small pieces of silver. However, this approach may result in a suboptimal solution if a valuable item, such as a small silver ring with significant archeological value that makes it the most valuable piece, is overlooked.

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

### Maximum Stock Profit

Given a list of stock prices for a given stock over time like `{50, 10, 4, 100, 1, 101, 5, 10}` return the maximum profit that can be made by buying and selling a single unit of this stock. Like `101 - 1 = 100` [Solution](max_stock_profit.go), [Tests](max_stock_profit_test.go)

### Activity Selector

Given a list of start and finish times for a few activities like `{1, 3, 0}` and `{4, 5, 6}` (first activity starts at 1 and ends at 4, second starts at 3 and ends at 5), return a maximal list of non-conflicting activities. [Solution](activity_selector.go), [Tests](activity_selector_test.go)
