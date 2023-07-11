# Complexity Analysis

Algorithms can be differentiated based on their time and space complexity. When evaluating an algorithm or operation, two crucial questions arise:

1. What is the time `t` required for execution?
2. How much memory space `s` does it utilize?

To address these questions, the Big O asymptotic notation, which characterizes how an algorithm performs with respect to time and space as the input size `n` increases, is employed.

## Big O

Big O is a mathematical notion commonly used to describe the impact on time or space as input size `n` increases. Seven Big O notations commonly used in algorithm complexity analysis are discussed in the following sections.

```ASCII
[Figure 1] Schematic diagram of Big O for common run times from fastest to slowest.

              O(1)                       O(Log n)                         O(n)
 ▲                              ▲                              ▲
 │                              │                              │
 │                              │                              │                      .
 │                              │                              │                    .
 │                              │                              │                  .
 │                              │                              │                .
 │                              │                              │              .
t│                             t│                      ...    t│            .
 │                              │              .........       │          .
 │                              │         ......               │        .
 │                              │     .....                    │      .
 │                              │   ...                        │    .
 │.........................     │ ..                           │  .
 │                              │.                             │.
 └────────────────────────►     └────────────────────────►     └────────────────────────►
              n                              n                              n

           O(n*Log n)                    O(Log 2^n)                       O(2^n)
 ▲                     .        ▲            .                 ▲        .
 │                    ..        │            .                 │        .
 │                    .         │            .                 │        .
 │                    .         │            .                 │        .
 │                    .         │           .                  │        .
 │                   .          │           .                  │       .
 │                  ..          │          .                   │       .
t│                 ..          t│          .                  t│      .
 │               ..             │         .                    │     .
 │              ..              │        .                     │    .
 │           ....               │       .                      │    .
 │       . . .                  │     ..                       │   .
 │   .....                      │  ...                         │  .
 │...                           │...                           │..
 └────────────────────────►     └────────────────────────►     └────────────────────────►
              n                              n                              n

              O(n!)
 ▲    .
 │    .
 │    .
 │    .
 │    .
 │    .
 │   .
t│   .
 │   .
 │   .
 │  .
 │ .
 │ .
 │.
 └────────────────────────►
              n
```

To understand the big O notation, let us focus on time complexity and specifically examine the O(n) diagram. This diagram depicts a decline in algorithm performance as input size increases. In contrast, the O(1) diagram represents an algorithm that consistently performs in constant time, with input size having no impact on its efficiency. Consequently, the latter algorithm generally outperforms the former.

However, it is essential to note that this is not always the case. In practice, a O(1) algorithm with a single time-consuming operation might be slower than a O(n) algorithm with multiple operations if the single operation in the first algorithm requires more time to complete than the collective operations in the second algorithm.

The Big O notation of an algorithm can be simplified using the following two rules:

1. Remove the constants. `O(n) + 2*O(n*Log n) + 3*O(K) + 5` is simplified to `O(n) + O(n*Log n) + O(K)`.
2. Remove non-dominant or slower terms. `O(n) + O(n*Log n) + O(K)` is simplified to `O(n*Log n)` because `O(n*Log n)` is the dominant term.

### Constant - O(K) or O(1)

Constant time complexity represents the most efficient scenario for an algorithm, where execution time remains constant regardless of input size. Achieving constant time complexity often involves eliminating loops and recursive calls. Examples:

* Reads and writes in a [hash table](./hashtable/README.md)
* Enqueue and Dequeue in a [queue](./queue/README.md)
* Push and Pop in a [stack](./stack/README.md)
* Find the minimum or maximum in a [heap](./heap/README.md)
* Remove the last element of a [doubly linked list](./linkedlist/README.md)
* [Max without conditions](./bit/max_function_without_conditions.go)

### Logarithmic - O(Log n)

Attaining logarithmic time complexity in an algorithm is highly desirable as it eliminates the need to iterate through every input to solve a given problem. Examples:

* Searching sorted items using [Binary Search](./dnc/binary_search.go)
* Inserting, Deleting, and Searching in a [Binary Search Tree](./tree/README.md)
* Push and Pop in a [heap](./heap/README.md)
* [Square Root](./dnc/square_root.go)
* [Median in a Stream](./heap/median_in_a_stream.go)

### Linear - O(n)

Linear time complexity is considered favorable when an algorithm traverses every input with no feasible way to avoid it. Examples:

* Remove the last element in a [singly linked list](./linkedlist/README.md)
* Search an unsorted [array](./array/README.md) or [linked list](./linkedlist/README.md)
* [Number of Islands](./graph/number_of_islands.go)
* [Missing Number](./hashtable/missing_number.go)

### O(n*Log n)

The time complexity of O(n*Log n) is commonly observed when it is necessary to iterate through all inputs and yield an outcome at the same time through an efficient operation. Sorting is a common example. It's impossible to sort items faster than O(n*Log n). Examples:

* [Merge Sort](./dnc/merge_sort.go) and [Heap Sort](./heap/README.md)
* [Knapsack](./greedy/knapsack.go)
* [Find Anagrams](./hashtable/find_anagrams.go)
* In order traversal of a [Binary Search Tree](./tree/README.md)

### Polynomial - O(n^2)

Polynomial time complexity marks the initial threshold of problematic time complexity for algorithms. This complexity often arises when an algorithm includes nested loops involving both an inner loop an outer loop. Examples:

* [Bubble Sort](./array/bubble_sort.go)
* [Cheapest Flight](./graph/cheapest_flights.go)
* [Remove Invalid Parentheses](./graph/remove_invalid_parentheses.go)

### Exponential O(2^n)

Exponential complexity is considered highly undesirable; however, it represents only the second-worst  complexity scenario. Examples:

* [Climbing Stairs](./recursion/climbing_stairs.go)
* [Towers of Hanoi](./dnc/towers_of_hanoi.go)
* [Generate Parentheses](./backtracking/generate_parentheses.go)
* Basic [Recursive](./recursion/README.md) implementation of Fibonacci

### Factorial O(n!)

Factorial time complexity represents the most severe time complexity for an algorithm. Understanding the scale of factorials is crucial, as even the estimated total number of atoms in the universe, which is approximately 10^80, is smaller than the factorial of 57. Example:

* [N queens](./backtracking/n_queens.go)
* [Permutations](./backtracking/permutations.go)
