# Divide-and-Conquer

The divide-and-conquer (DNC) paradigm is a common approach to solving problems [recursively](../recursion). It involves breaking down a problem into smaller and smaller sub-problems until a base case is reached, where the problem can be directly solved. The divide-and-conquer method consists of three steps:

1. Divide: The problem is divided into smaller sub-problems, typically by partitioning the input data into two or more subsets.
2. Conquer: The smaller sub-problems are solved recursively using the same algorithm, typically by applying the divide-and-conquer approach again.
3. Combine: The solutions to the smaller sub-problems are combined to solve the original problem.

## Implementation

The binary search algorithm is a classic Divide-and-Conquer algorithm used to find a specific value in a sorted list. The search process begins by comparing the target value to the middle element of the list. If they are not equal, the half of the list in which the target value cannot be present is eliminated. This process is repeated on the remaining half of the list until the target value is found or until there are no more elements to search. Both recursive and iterative implementations of binary search are shown below.

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

Go comes with a built-in binary search function that returns the index of the first element satisfying a given condition. This function eliminates the need to manually implement the binary search algorithm.

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

If used inappropriately, DNC algorithms can lead to an exponential number of unnecessary recursive calls, resulting in a time complexity of O(2^n). However, DNC algorithms can be very effective if an appropriate dividing strategy and a base case that can be solved directly are identified. They have a time complexity as low as O(Log n), such as in the case of binary search. As DNC algorithms are recursive, their complexity analysis is similar to [recursive](../recursion) algorithms.

## Application

DNC algorithms are suitable for solving problems that can be divided into smaller sub-problems that can be solved similarly. Examples of such problems include sorting algorithms (e.g., merge sort, quicksort), searching algorithms (e.g., binary search), and matrix multiplication.

## Rehearsal

* [Binary Search](binary_search_test.go), [Solution](binary_search.go)
* [Square Root with Binary Search](square_root_test.go), [Solution](square_root.go)
* [Rate Limit](rate_limit_test.go), [Solution](rate_limit.go)
* [Towers of Hanoi](towers_of_hanoi_test.go), [Solution](towers_of_hanoi.go)
* [Merge Sort](merge_sort_test.go), [Solution](merge_sort.go)
* [Quick Sort](quick_sort_test.go), [Solution](quick_sort.go)
