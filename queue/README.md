# Queue
Queues are data structures that operate on the principle of First-In-First-Out (FIFO), enabling elements to be added to the rear of the queue through the enqueue operation and removed from the front of the queue through the dequeue operation.

In the real world queues are formed in situations where a first-come, first-served service is provided, such as at toll booths on highways.

Another variation of queues is the double-ended queue, which allows for dequeuing from both ends, thereby facilitating both FIFO and LIFO (Last In, First Out) data structures.

## Implementation

Similar to [stacks](../stack) queues can be implemented either with arrays or linked list.Here's a slice implementation of an integer queue:

```Go
package main

import (
	"errors"
	"fmt"
)

var queue []int

func main() {
	enqueue(1) // [1]
	enqueue(2) // [1,2]
	enqueue(3) // [1,2,3]
	dequeue()  // [2,3]
	dequeue()  // [3]
	dequeue()  // []
}

func enqueue(val int) {
	queue = append(queue, val)
}

func dequeue() (int, error) {
	if len(queue) == 0 {
		return -1, errors.New("Queue is empty")
	}
	tmp := queue[0]
	queue = queue[1:len(queue)]
	return tmp, nil
}
```

Here's a [linked list](../linkedlists) implementation with the same exact outcome.

```Go
package main

import (
	"container/list"
	"errors"
)

var queue = list.New()

func main() {
	enqueue(1) // [1]
	enqueue(2) // [1,2]
	enqueue(3) // [1,2,3]
	dequeue()  // [2,3]
	dequeue()  // [3]
	dequeue()  // []
}

func enqueue(val int) {
	queue.PushBack(val)
}

func dequeue() (int, error) {
	if queue.Len() == 0 {
		return -1, errors.New("Queue is empty")
	}
	tmp := queue.Front().Value.(int)
	queue.Remove(queue.Front())
	return tmp, nil
}
```

## Complexity

Enqueue and dequeue operations are both done in O(1) times.

## Application

Queues are widely utilized in solving graph-related problems and managing capacity in various contexts such as printers, where tasks are processed in the order of their arrival. They are also utilized in situations where a first-come-first-serve approach is necessary.

## Rehearsal

### A Queue Using Stacks

Implement a queue of integers using stacks. [Solution](queue_using_stacks.go), [Tests](queue_using_stacks_test.go)

### Implement a Queue Using a Circular Array

Implement a queue of integers using stacks. [Solution](queue_using_circular_array.go), [Tests](queue_using_circular_array_test.go)

### Is Binary Tree Symmetrical

```ASCII
      2
    /   \
   /     \
  4       4
 / \     / \
5   6   6   5
```

Given a binary tree the one above "2,4,4,5,6,6,5", true if it is symmetric (mirrored from the root), and false otherwise. [Solution](is_tree_symmetrical.go), [Tests](is_tree_symmetrical_test.go)

## Generate Binary Numbers

Given a number n (n>=0) count from 0 to n in binary. [Solution](generate_binary_numbers.go), [Tests](generate_binary_numbers_test.go)

### Find The Maximum of Sub-array of Length K

Given a set an array of numbers like {1,2,3,4,5} and a number k like 2 return the maximum in each k-lengthed sub array. e.g. {2,3,4,5} corresponding to the max in each set of the sub arrays {{1,2},{2,3},{3,4},{4,5}}. [Solution](maximum_of_sub_arrays.go), [Tests](maximum_of_sub_arrays_test.go)
