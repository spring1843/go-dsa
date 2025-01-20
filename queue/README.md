# Queue

Queues are data structures that operate on the principle of First-In-First-Out (FIFO), enabling elements to be added to the rear of the queue through the enqueue operation and removed from the front of the queue through the dequeue operation.

In the real world, queues are formed when a first-come, first-served service is provided, such as at highway toll booths.

Another variation of queues is the double-ended queue, which allows for dequeuing from both ends, facilitating both FIFO and LIFO (Last In, First Out) data structures.

## Implementation

Like [stacks](../stack), queues can be implemented using doubly [linked lists](../linkedlist/) or [arrays and slices](../array/). Here is a linked list implementation:

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

Enqueue and dequeue operations both perform in O(1) times in the linked list implementation. In other traditional languages like C the linked list approach is considered to be faster than the slice approach because both enqueue and dequeue operations are O(n) due to the slice size change. In Go however slices are managed intelligently behind the scene and perform very well for just about all purposes.

We can use Go's built-in benchmarking tooling to see which implementation is faster. This is done in [slice_vs_linked_list_bench_test.go](./slice_vs_linked_list_bench_test.go). It can be executed by running `go test -bench=. -test.benchmem` in this directory. The output shows that the slice implementation is almost seven times faster.

```Shell
pkg: github.com/spring1843/go-dsa/queue
BenchmarkLinkedListQueue-8      17956176                83.73 ns/op           56 B/op          1 allocs/op
BenchmarkSliceQueue-8           100000000               11.79 ns/op           45 B/op          0 allocs/op
PASS
ok      github.com/spring1843/go-dsa/queue      3.775s
```

## Application

Queues are widely utilized in solving graph-related problems and managing capacity in various contexts, such as printers, where tasks are processed in the order of arrival. They are also utilized in situations where a first-come-first-serve approach is necessary.

## Rehearsal

* [Queue Using Stacks](./queue_using_stacks_test.go), [Solution](./queue_using_stacks.go)
* [Circular Queue Array](./circular_queue_using_array_test.go), [Solution](./circular_queue_using_array.go)
* [Symmetrical Binary Tree](./symmetrical_binary_tree_test.go), [Solution](./symmetrical_binary_tree.go)
* [Generate Binary Numbers](./generate_binary_numbers_test.go), [Solution](./generate_binary_numbers.go)
* [Max Sub-array of size K](./max_of_sub_arrays_test.go), [Solution](./max_of_sub_arrays.go)
* [String Permutations](./string_permutations_test.go), [Solution](./string_permutations.go)
