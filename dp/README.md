# Dynamic Programming

Dynamic Programming (DP) and [divide-and-conquer](../dnc) share a common strategy of breaking down a problem into smaller sub-problems. However, DP achieves superior algorithmic performance by solving each sub-problem only once and preemptively eliminating sub-problems that cannot contribute to optimal solutions. This approach enables DP to avoid the redundancy inherent in DNC and produce more efficient solutions.

## Implementation

Dynamic Programming (DP) approach to algorithms typically involves four steps:

1. Characterize an optimal solution.
2. Define the value of an optimal solution through [recursion](../recursion).
3. Compute the value of the optimal solution in a bottom-up manner.
4. Determine the optimal solution using the values computed in the previous steps.

If only the solution value is needed, step 4 may be omitted. Conversely, when the solution is required, it is advisable to consider the necessary values at the final step. This will facilitate the storage of pertinent information at step 3 and simplify step 4.

There are two general methods for writing DP algorithms: top-down and bottom-up approaches.

In the top-down approach, a caching mechanism is utilized to store each sub-problem solution and prevent redundancy. This technique is also known as memoization.

In the bottom-up approach, sub-problems are solved in order of size, with the smaller ones tackled first. Since all subsequent smaller sub-problems are already solved when addressing a particular sub-problem, the result is stored.

## Complexity

Top-down and bottom-up approaches usually perform similarly.

## Application

DP is well-suited for tackling an array of complex problems, including logistics, game theory, machine learning, resource allocation, and investment policies. In graph theory, DP is commonly used to determine the shortest path between two points. DP algorithms are particularly effective at optimizing problems with a vast number of potential solutions, where the goal is to identify the most optimal one.

## Rehearsal

* [Rod Cutting](rod_cutting_test.go), [Solution](rod_cutting.go)
* [Sum Up to Number](sum_up_to_integer_test.go), [Solution](sum_up_to_integer.go)
* [House Robber](house_robber_test.go), [Solution](house_robber.go).
* [Minimum Deletion to Make a Palindrome](minimum_deletion_to_make_palindrome_test.go), [Solution](minimum_deletion_to_make_palindrome.go)
* [Word Distance](word_distance_test.go), [Solution](word_distance.go.go)
