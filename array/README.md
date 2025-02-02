# Array

Arrays are a fundamental and essential data structure in computer science. They consist of a fixed-size contiguous memory block and offer O(1) read and write time complexity. In Go Like most other programming languages arrays are a built-in data type.

To provide a real-world analogy, consider an array of athletes preparing for a sprinting match. Each athlete occupies a specific position within the array, typically denoted as 1, 2,…, n. While it is technically possible for each athlete to be in a different position, the positions generally carry some form of significance, such as alphabetical order or seniority within the sport.

## Implementation

In the Go programming language, [arrays](https://go.dev/blog/slices) are considered values rather than pointers and represent the entirety of the array. Whenever an array is passed to a function, a copy is created, resulting in additional memory usage. To avoid this, it is possible to pass a pointer to an array, or use slices instead. The size of the array is constant and it must be known at compile time, and there is no need to use the built-in `make` function when defining arrays.

```Go
package main

import "fmt"

func main() {
	var nums1 [2]int
	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums1, nums2) // Prints [0 0] [1 2 3]
}
```

Although arrays are fundamental data structures in Go, their constant size can make them inflexible and difficult to use in situations where a variable size is required. To address this issue, Go provides [slices](https://blog.golang.org/slices-intro), an abstraction of arrays that offer more convenient access to sequential data typically stored in arrays. When a slice is passed to a function, only the slice header is passed, but it still references the same underlying array data, hence it is possible for the callee to modify the values of the slice and send them back to the caller.

Slices enable adding values using the `append` function, allowing dynamic resizing. Additionally, selectors of the format [low:high] can be used to select or manipulate data in the slice. By utilizing slices instead of arrays, Go programmers benefit from a more flexible and powerful tool to manage their data structures.

```Go
package main

import "fmt"

func main() {
	// Create a slice with elements 1, 2, 3.
	nums := []int{1, 2, 3}
	nums = append([]int{0}, nums...) // Add a new element to the start
	nums = append(nums, 4)           // Add a new element to the end
	nums = nums[:len(nums)-1]        // Removes the last element
	fmt.Println(nums)                // Prints [0 1 2 3]
}
```

The [make](https://golang.org/pkg/builtin/#make) function can create a zeroed slice of a given length and capacity.

```Go
package main

import "fmt"

func main() {
	nums := make([]int, 2)
	nums[0], nums[1] = 0, 1
	fmt.Println(nums, len(nums), cap(nums)) // Prints [0 1] 2 2
}
```

Slice expressions in the form of input[low:high] can be used to manipulate slices or access their elements

```Go
package main

import "fmt"

func main() {
    // Initialize a slice with 6 elements.
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = nums[:len(nums)-1] // Drop the last element
	nums = nums[1:]           // Drop the first element
	nums = nums[1:]           // Keep all elements from index 1 to the end
	nums = nums[:2]           // Keep all elements up to (but not including) index 2
	nums = nums[1:2]          // Keep only the element at index 1
	fmt.Println(nums)         // Prints [4]
}
```

## Complexity

Accessing an element within an array using an index has O(1) time complexity. This means that regardless of the size of the array, read and write operations for a given element can be performed in constant time.

While arrays are useful for certain tasks, searching an unsorted array can be a time-consuming O(n) operation. Since the target item could be located anywhere in the array, every element must be checked until the item is found. Due to this limitation, alternative data structures such as [linked-lists](../linkedlist/), [trees](../tree/) and [hash tables](../hashtable/) are often more suitable for search operations.

Addition and deletion operations are O(n) operations in Arrays. Removing an element can create an empty slot that must be eliminated by shifting the remaining items. Similarly, adding items to an array may require shifting existing items to create space for the added item. These inefficiencies can make alternative data structures, such as [trees](../tree) or [hash tables](../hashtable), more suitable for managing operations involving additions and deletions.

## Application

Arrays are used wherever sequential data or more than one piece of data is needed. The fast read and write access to a given element makes arrays suitable for implementing over other data structures such as [strings](../strings), [stacks](../stack), [queues](../queue), and [hash tables](../hashtable).

## Rehearsal

* [Reverse Array In-place](./reverse_inplace_test.go), [Solution](./reverse_inplace.go)
* [Add Two Numbers Represented as Slices](./add_slice_of_numbers_test.go), [Solution](./add_slice_of_numbers.go)
* [Find Duplicate in Array](./find_duplicate_in_array_test.go), [Solution](./find_duplicate_in_array.go)
* [Zero Sum Triplets](./zero_sum_triplets_test.go), [Solution](./zero_sum_triplets.go)
* [Product of All Other Elements](./product_of_all_other_elements_test.go), [Solution](./product_of_all_other_elements.go)
* [Equal Sum Sub-arrays](./equal_sum_subarrays_test.go), [Solution](./equal_sum_subarrays.go)
* [Rotate K Times](./rotate_k_steps_test.go), [Solution](./rotate_k_steps.go)
* [Bubble Sort](./bubble_sort_test.go), [Solution](./bubble_sort.go)
* [Insertion Sort](./insertion_sort_test.go), [Solution](./insertion_sort.go)
