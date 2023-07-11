# Queue

Queues are data structures that operate on the principle of First-In-First-Out (FIFO), enabling elements to be added to the rear of the queue through the enqueue operation and removed from the front of the queue through the dequeue operation.

In the real world, queues are formed in situations where a first-come, first-served service is provided, such as at toll booths on highways.

Another variation of queues is the double-ended queue, which allows for dequeuing from both ends, thereby facilitating both FIFO and LIFO (Last In, First Out) data structures.

## Implementation

Similar to [stacks](../stack), queues can be implemented using doubly [linked lists](../linkedlist/) or [arrays and slices](../array/). Here is a linked list implementation:

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
		return -1, errors.New("queue is empty")
	}
	return queue.Remove(queue.Front()).(int), nil
}
```

Here's a slice implementation of an integer queue:

```Go
package main

import "errors"

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
		return -1, errors.New("queue is empty")
	}
	tmp := queue[0]
	queue = queue[1:len(queue)]
	return tmp, nil
}
```

## Complexity

Enqueue and dequeue operations both perform in O(1) times.

## Application

Queues are widely utilized in solving graph-related problems and managing capacity in various contexts, such as printers, where tasks are processed in the order of their arrival. They are also utilized in situations where a first-come-first-serve approach is necessary.

## Rehearsal

* [A Queue Using Stacks](queue_using_stacks_test.go), [Solution](queue_using_stacks.go)
* [Implement a Queue Using a Circular Array](queue_using_circular_array_test.go), [Solution](queue_using_circular_array.go)
* [Is Binary Tree Symmetrical](is_tree_symmetrical_test.go), [Solution](is_tree_symmetrical.go)
* [Generate Binary Numbers](generate_binary_numbers_test.go), [Solution](generate_binary_numbers.go)
* [Find The Maximum Sub-array of Length K](maximum_of_sub_arrays_test.go), [Solution](maximum_of_sub_arrays.go)
