# Dynamic Programming

Dynamic Programming (DP) and [divide-and-conquer](../dnc) share a common strategy of breaking down a problem into smaller subproblems. However, DP achieves superior algorithmic performance by solving each subproblem only once and preemptively eliminating subproblems that cannot contribute to optimal solutions. This approach enables DP to avoid the redundancy inherent in DNC and produce more efficient solutions.

## Implementation

A Dynamic Programming (DP) approach to algorithms typically entails four steps:

1. Characterizing an optimal solution
2. Defining the value of an optimal solution through [recursion](../recursion)
3. Computing the value of the optimal solution in a bottom-up manner
4. Determining the optimal solution using the values computed in the previous steps.

If only the value of the solution is needed, step 4 may be omitted. Conversely, when the solution is required, it is advisable to consider the necessary values at the final step to facilitate storage of pertinent information at step 3 and simplify step 4.

There are two general methods for writing DP algorithms: the top-down and bottom-up approaches.

In the top-down approach, a caching mechanism is utilized to store each subproblem solution and prevent redundancy. This technique is also known as memoization.

In the bottom-up approach, subproblems are solved in order of size, with smaller ones tackled first. Since all subsequent smaller subproblems are already solved when addressing a particular subproblem, the result is stored.

## Complexity

The top-down and bottom-up approaches typically perform similarly.

## Application

DP is well-suited for tackling an array of complex problems, including those in logistics, game theory, machine learning, resource allocation, and investment policies. In graph theory, DP is commonly used to determine the shortest path between two points. DP algorithms are particularly effective in optimizing problems that have a vast number of potential solutions, where the goal is to identify the most optimal one.

## Rehearsal

* [Rod Cutting](rod_cutting_test.go), [Solution](rod_cutting.go)
* [Sum Up to Number](sum_up_to_integer_test.go), [Solution](sum_up_to_integer.go)
* [House Robber](house_robber_test.go), [Solution](house_robber.go).
* [Minimum Deletion to Make a Palindrome](minimum_deletion_to_make_palindrome_test.go), [Solution](minimum_deletion_to_make_palindrome.go)
* [Word Distance](word_distance_test.go), [Solution](word_distance.go.go)
