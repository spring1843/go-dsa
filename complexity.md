# Complexity Analysis

Algorithms can be differentiated based on their time and space complexity. When evaluating an algorithm or operation, two crucial questions arise:

1. What is the time `t` required for execution?
2. How much memory space `s` does it utilize?

To address these questions, the Big O asymptotic notation is employed, which characterizes how an algorithm performs with respect to time and space as the input size `n` increases.

## Big O

Big O is mathematical notion commonly used to describe and compare the worst case performance of an algorithm in both time and space given inputs of size `n`. Seven Big O notations that commonly used in algorithm complexity analysis are discussed in the following sections.

```ASCII
[Figure 1] Schematic diagrams of Big O for common run times from fastest to slowest.

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

           O(N Log N)                    O(Log 2^n)                       O(2^n)
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

To understand the big O notation, let us focus on time complexity and specifically examine the O(n) diagram. This diagram depicts a decline in the algorithm's performance as the input size increases. In contrast, the O(1) diagram represents an algorithm that consistently performs in constant time, with the input size having no impact on its efficiency. Consequently, the latter algorithm generally outperforms the former.

However, it is essential to note that this is not always the case. In practice, a O(1) algorithm with a single time-consuming operation might be slower than a O(n) algorithm with multiple operations if the single operation in the first algorithm requires more time to complete than the collective operations in the second algorithm.

Big O notation of an algorithm can be simplified using the following two rules:

1. Remove constants. `O(n) + 2*O(n Log n) + 3*O(K) + 5` is simplified to `O(n) + O(n Log n) + O(K)`.
2. Remove non dominant, or slower terms. `O(n) + O(n Log n) + O(K)` is simplified to `O(n Log n)` because `O(n Log n)` is the most dominant term..

### Constant - O(K) or O(1)

Constant time complexity represents the most efficient scenario for an algorithm, where the execution time remains constant regardless of the input size. Achieving constant time complexity often involves eliminating loops and recursive calls. Examples:

* Reads and writes in a [hash table](../hashtable)
* Enqueue and Dequeue in a [queue](../queue)
* Push and Pop in a [stack](../stack)
* Finding the minimum or maximum in [heap](../heap)
* Removing the last element of a [doubly linked list](../linkedlist)

### Logarithmic - O(Log n)

Attaining logarithmic time complexity in an algorithm is highly desirable as it eliminates the need to iterate through every input in order to solve a given problem. Examples:

* Searching sorted items using [Binary Search](../dnc)
* Inserting, Deleting and Searching in a [Binary Search Tree](../tree)
* Push and Pop in [heap](../heap)

### Linear - O(n)

Linear time complexity is considered favorable when an algorithm necessitates traversing every input, with no feasible way to avoid it. Examples:

* Removing the last element in a [singly linked list](../linkedlist)
* Searching an unsorted [array](../array) or [linked list](../linklist)

### O(n Log n)

The time complexity of O(n log n) is commonly observed when it is necessary to iterate through all inputs, and can yield an out come at the same time through an efficient operation. Sorting is a common example. It's not possible to sort items faster than O(n log n). Examples:

* [Merge Sort](../dnc) and [Heap Sort](../heap)
* In order traversal of a [Binary Search Tree](../tree)

### Polynomial - O(n^2)

Polynomial time complexity marks the initial threshold of problematic time complexity for algorithms. This complexity often arises when an algorithm includes nested loops, involving both an inner loop and an outer loop. Examples:

* Bubble sort rehearsal problem in [array](../array)
* Naive way of searching an unsorted [array](../array) for duplicates by using nested loops

### Exponential O(2^n)

Exponential complexity is considered highly undesirable; however, it represents only the second-worst  complexity scenario. Examples:

* Basic [Recursive](../recursion) implementation of Fibonacci
* Tower of Hanoi rehearsal in [divide and conquer](../dnc)

### Factorial O(n!)

Factorial time complexity represents the most severe time complexity for an algorithm. Understanding the scale of factorials is crucial, as even the estimated total number of atoms in the universe, which is approximately 10^80, is smaller than the factorial of 57. Example:

* Permutations rehearsal in [back tracking](../backtracking)
