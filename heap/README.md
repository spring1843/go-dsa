# Heap

Heaps are tree data structures that retain the minimum or maximum of the elements pushed into them. There are two types of heaps: minimum and maximum heaps.

A heap must satisfy two conditions:

1. The structure property requires that the heap be a complete binary search [tree](../tree), where each level is filled from left to right, and all levels except the bottom are full.
2. The heap property requires that the children of a node be larger than or equal to the parent node in a min heap and smaller than or equal to the parent in a max heap, meaning that the root is the minimum in a min heap and the maximum in a max heap.

As a result, if all elements are pushed to the min or max heap and then popped one by one, a sorted list in ascending or descending order is attained. This sorting technique known as [heap sort](./heap_sort_test.go) works in O(n*Logn) time. While other comparison based sorting algorithms are available, none can be more efficient than O(n*Log n).

When pushing an element to a heap, because of the structure property, the new element is always added to the first available position on the lowest level of the heap, filling from left to right. Then to maintain the heap property, if the newly inserted element is smaller than its parent in a min heap (larger in a max heap), the newly added element is percolated up by being swapped with its parent. The child and parents are swapped until the heap property is achieved.

```ASCII
[Figure 1] Minimum heap push operation

  20	    20		  15			 15		15		5
 /  \	   /  \		 /  \			/  \	      	/  \	      /   \
25  30	  25  30  	20  30	               20  30	      5   30	    15     30
         /	       /      	              /  \	     / \           /  \
        15	      25	             25   5	   25  20	 25    20
										
	(A) Add 15		     	    (B) Add 5
```

The pop operation in a heap starts by replacing the root with the rightmost leaf. Then the root is swapped with the smaller child in a min heap (and the larger child in a max heap). The root is removed and the new root is percolated down until the heap property is achieved.

```ASCII
[Figure 2] Minimum heap pop operation

       5	    20	       15		  25	   20		  30
     /   \	   /  \	      /  \		 /  \	  /  \		 /
    15   30	  15  30     20  30		20  30   25  30		25
   /  \		 /	    /
  25  20	25	   25

	(A) Remove				(B) Remove		(C) Remove
```

An example implementation of this is provided as a [solution](./heap_sort.go) to the [heap sort](./heap_sort_test.go) rehearsal.

## Implementation

The Go standard library includes an implementation of heap in [container/heap](https://golang.org/pkg/container/heap/). Below is an example of a maximum heap implementation:

```Go
package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func main() {
	h := new(maxHeap)
	for i := 1; i <= 10; i++ {
		heap.Push(h, i)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(heap.Pop(h))
	}

}

func (m maxHeap) Len() int            { return len(m) }
func (m maxHeap) Less(i, j int) bool  { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x interface{}) { *m = append(*m, x.(int)) }

func (m *maxHeap) Pop() interface{} {
	old := *m
	tmp := old[len(old)-1]
	*m = old[0 : len(old)-1]
	return tmp
}
```

To utilize a heap to store a particular type, certain methods such as len and less must be implemented for that type to conform to the heap interface. By default, the heap is a min heap, with each node smaller than its children. However, the package provides the flexibility to define what "being less than" means. For instance, changing `m[i] > m[j]` to `m[i] < m[j]` would transform the heap into a minimum heap.

In Go, heaps are implemented with slices. The heap property is maintained such that the left child of the node at index `i` (where i is greater than or equal to 1) is always located at `2i`, and the right child is at `2i+1`. If the slice already contains elements before pushing, the heap must be initialized using `heap.Init(h Interface)` to establish the order.

## Complexity

Pushing and popping are O(LogN) operations in heaps. On the other hand, initializing a heap, which involves pushing N elements, has a time complexity of O(n*Log n).

## Application

Heaps in the form of priority queues are used for scheduling in operating systems and job schedulers. They are also used in simulating and implementing scheduling and priority systems for capacity management in hospitals and businesses.

Priority queues implemented as heaps are commonly used in job scheduling, for example scheduling the execution of different processes in operating systems. They are also employed in simulating and implementing priority and scheduling systems to manage capacity in hospitals and businesses.

## Rehearsal

* [Heap Sort](./heap_sort_test.go), [Solution](./heap_sort.go)
* [Kth Largest Element](./kth_largest_element_test.go), [Solution](./kth_largest_element.go)
* [Merge Sorted Lists](./merge_sorted_list_test.go), [Solution](./merge_sorted_list.go)
* [Median in a Stream](./median_in_a_stream_test.go), [Solution](./median_in_a_stream_test.go)
* [Regular Numbers](./regular_numbers_test.go), [Solution](./regular_numbers_test.go)
* [Kth Closest Points to the Center](./k_closest_points_to_origin_test.go), [Solution](./k_closest_points_to_origin.go)
* [Sliding Maximum](./sliding_maximum_test.go), [Solution](./sliding_maximum.go)
