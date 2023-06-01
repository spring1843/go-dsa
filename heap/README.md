# Heap

Heaps are tree data structures that function by retaining the minimum or maximum of the elements inserted into them. There are two types of heaps: minimum heaps and maximum heaps.

A heap must satisfy two conditions:

1. The structure property requires that the heap be a complete binary search [tree](../tree), where each level is filled left to right, and all levels except the bottom are full.
2. The heap property requires that the children of a node be larger than or equal to the parent node in a min heap and smaller than or equal to the parent in a max heap, meaning that the root is the minimum in a min heap and the maximum in a max heap.

As a result, if you push elements to the min or max heap and then pop them one by one, you will obtain a list that is sorted in ascending or descending order, respectively. This sorting technique is also an O(NLogN) algorithm known as heap sort. Although there are other sorting algorithms available, none of them are faster than O(NLogN).

When pushing a new element to a heap, because of the structure property we always add the new element to the first available position on the lowest level of the heap, filling from left to right. Then to maintain the heap property, if the newly inserted element is smaller than its parent in a min heap (larger in a max heap), then we swap it with its parent. We continue swapping the  swapped element with its parent until the heap property is achieved.

```ASCII
[Figure 1] Minimum heap push operation

  20	    20		  15				 15			   15			 5
 /  \	   /  \		 /  \				/  \		  /  \		   /   \
25  30	  25  30  	20  30			   20  30		 5   30		  15    30
		 /		   /				  /  \		    /  \		 /  \
		15		  25				 25   5		   25   20		25   20
										
	(A) Add 15							(B) Add 5
```

The pop operation in a heap starts by replacing the root with the right most leaf. Then we swap the root element down with the smaller child in a min heap (and larger child in a max heap) until the heap property is achieved.

```ASCII
[Figure 2] Minimum heap pop operation

       5		20		15			  25	   20			30
     /   \	   /  \	   /  \			 /  \	  /  \			/
	15   30	  15  30  20  30		20  30   25  30		   25
   /  \		 /		 /
  25  20	25		25

	(A) Remove						(B) Remove			(C) Remove
```

## Implementation

The Go standard library includes an implementation of a heap in [container/heap](https://golang.org/pkg/container/heap/). Below is an example of a maximum heap implementation:

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

To utilize a heap to store a particular type, certain methods such as len and less must be implemented for that type to conform to the heap interface. By default, the heap is a min heap, where each node is less than its children. However, the package provides the flexibility to define what "being less than" means. For instance, changing `m[i] > m[j]` to `m[i] < m[j]` would transform the heap into a minimum heap.

In Go, the heap implementation is based on slices. The heap property is maintained such that the left child of the node at index `i` (where i is greater than or equal to 1) is always located at `2i`, and the right child is at `2i+1`. If the slice already contains elements before any pushing operation, the heap must be initialized using heap.Init(h Interface) to establish the order.

## Complexity

The time complexity of pushing and popping heap elements is O(LogN). On the other hand, initializing a heap, which involves pushing N elements, has a time complexity of O(NLogN).

The insertion strategy entails percolating the new element up the heap until the correct location is identified. Similarly, the deletion strategy involves percolating down the heap.

Pushing and Popping heap elements are all O(LogN) operations. The strategy for inserting is the new element is percolating up the heap until the correct location is found. similarly the strategy for deletion is to percolate down.

## Application

Heaps in the form of priority queues are used for scheduling in operating systems and job schedulers. They are also used in simulating and implementing scheduling and priority systems for capacity management in hospitals and businesses.

Priority queues implemented as heaps are utilized in job scheduling, for example scheduling the execution of different processes in an operating systems. They are also employed in simulating and implementing priority and scheduling systems to manage capacity in hospitals and businesses.

## Rehearsal

### Kth Largest Element

Given an array of integers like {3,5,6,3,1,2,9} and k an integer like 3 return the kth largest element like 5. [Solution](kth_largest_element.go), [Test](kth_largest_element_test.go)

### Merge Sorted Lists

Given multiple sorted linked lists like {1->2, 1->3->4, 4->5}, join them into one like 1->1->2->3->4->4->5. [Solution](merge_sorted_list.go), [Test](merge_sorted_list_test.go)

### Median in a Stream

Given a stream of integers like {1,2,3,4}, output the median after each input like {1, 1.5, 2, 2.5}. [Solution](median_in_a_stream_test.go), [Test](median_in_a_stream_test.go)

### Kth Closest Points to the Center

Given coordination of a point on an x,y axis and an integer k, return k closest points to the origin. [Solution](k_closest_points_to_origin.go), [Test](k_closest_points_to_origin_test.go)

### Sliding Maximum

Given a list of integers like `{1, 4, 5, -2, 4, 6}`, and a positive integer k like 3, return the maximum of each slice of the array when a window of length k is moved from left to the right in the array like `{4}`. [Solution](sliding_maximum.go), [Test](sliding_maximum_test.go)
