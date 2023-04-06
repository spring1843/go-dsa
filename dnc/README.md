# Divide-and-Conquer

The divide-and-conquer (DNC) paradigm is a common approach to solving problems [recursively](../recursion). It involves breaking down a problem into smaller and smaller sub-problems until a base case is reached, where the problem can be solved directly. The divide-and-conquer method consists of three steps:

1. Divide: The problem is divided into smaller sub-problems, typically by partitioning the input data into two or more subsets.
2. Conquer: The smaller sub-problems are solved recursively using the same algorithm, typically by applying the divide-and-conquer approach again.
3. Combine: The solutions to the smaller sub-problems are combined to obtain a solution to the original problem.

## Implementation

The binary search algorithm is a classic example of a Divide-and-Conquer algorithm that is commonly used to find a specific value in a sorted list or array. The search process begins by comparing the target value to the middle element of the list. If they are not equal, the half of the list where the target value cannot be present is eliminated. This process is repeated on the remaining half of the list until the target value is found or until there are no more elements to search. Binary search can be implemented iteratively or recursively, and both implementations are shown below.

```Go
package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearchRecursive(list, 0, len(list), 5)) // Prints 4
	fmt.Println(binarySearchIterative(list, 5))               // Prints 4
}

func binarySearchRecursive(list []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if list[mid] == target {
		return mid
	} else if list[mid] > target {
		return binarySearchRecursive(list, low, mid-1, target)
	} else {
		return binarySearchRecursive(list, mid+1, high, target)
	}
}

func binarySearchIterative(list []int, target int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] < target {
			low = mid + 1
		} else if list[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
```

A pre-implemented binary search function is available in Go which returns the index of the first element satisfying a given condition. This function eliminates the need to manually implement the binary search algorithm.

```Go
package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(search(list, 5)) // Prints 4
	fmt.Println(search(list, 6)) // Prints -1
}

func search(list []int, target int) int {
	i := sort.Search(len(list), func(i int) bool { return list[i] >= target })
	if i < len(list) && list[i] == target {
		return i
	}
	return -1
}
```

## Complexity

If used inappropriately, DNC algorithms can lead to an exponential number of unnecessary recursive calls, resulting in a time complexity of O(2^n). However, if an appropriate dividing strategy and base case that can be solved directly are identified, DNC algorithms can be very effective, with a time complexity as low as O(log n) in the case of binary search. As DNC algorithms are recursive in nature, their complexity analysis is analogous to that of [recursive](../recursion) algorithms.

## Application

DNC algorithms are suitable for solving problems that can be divided into smaller sub-problems that can be solved in a similar way. Examples of such problems include sorting algorithms (e.g., merge sort, quicksort), searching algorithms (e.g., binary search), and matrix multiplication.

## Rehearsal

### Binary Search

Given a sorted set of integers like `{1,2,3,4,6}`, and a target int like `4` find its position in the set like `3`. [Solution](binary_search.go) [Test](binary_search_test.go)

### Square Root with Binary Search

Given a number and precision, return the square root of the number using the binary search algorithm. [Solution](square_root.go) [Test](square_root_test.go)

### Rate Limit

Given a number of allowed calls per second, write an IsAllowed function which returns false if the call should be rate limited and true if the call should be allowed (i.e. is within the range of allowed calls per second). [Solution](rate_limit.go) [Test](rate_limit_test.go)

### Towers of Hanoi

Given n, number of disks, and start and end tower, return the moves it takes to move all disks from start to end tower. The disks are stacked on top of each other with the lightest being on top and heaviest being in the bottom. A heavier disk cannot be placed on a lighter disk. You can move one disk at a time. [Solution](towers_of_hanoi.go) [Test](towers_of_hanoi_test.go)

### Merge Sort

Given a list of integers like `{3,1,2}`, return a sorted set like `{1,2,3}` using Merge Sort. [Solution](merge_sort.go) [Test](merge_sort_test.go)
