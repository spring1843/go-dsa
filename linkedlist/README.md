# Linked List

Linked lists are a collection of nodes, each capable of storing at least one data element and linked to the next node via a reference. One of the key advantages of linked lists over [arrays](../array) is their dynamic size, which allows for items to be added or removed without necessitating the resizing or shifting of other elements.

Two types of linked lists exist: singly linked lists, in which each node is linked only to the next node, and doubly linked lists, in which each node is connected to both the next and previous nodes.

## Implementation

When implementing linked lists, two key concepts come into play: the linked list itself, which includes references to the head and tail nodes, and the individual nodes which hold:

* At least one piece of data, such as name or number
* A reference to the next node
* A reference to the previous node in doubly linked lists

A node `struct` can be defined to create a linked list in Go, and a pointer to the same structure can be used to link the node to the next one. For instance, a singly linked list with the data `1->2->3->4->5` can be created using the following example:

```Go
package main

type node struct {
	val  int
	next *node
}

var head *node

func main() {
	for i := 5; i > 0; i-- {
		addToFront(newNode(i))
	}
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

func addToFront(node *node) {
	if head != nil {
		node.next = head
	}
	head = node
}
```

A pointer is typically stored for the first and sometimes the last items in a singly linked list. Adding items to the front or back of the list is a constant-time operation. However, deleting the last item can be challenging, as the last item's pointer needs to be updated to the second-to-last item. This is where referencing the last item in each node proves useful. In contrast, doubly linked lists maintain pointers to the previous and next nodes, making deletion operations less expensive.

The Go standard library contains an implementation of [doubly linked lists](https://golang.org/pkg/container/list/). In the following example, numbers from 1 to 10 are added to the list. Even numbers are removed, and the resulting linked list containing odd numbers is printed.

```Go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	numbers := list.New()
	for i := 1; i <= 10; i++ {
		numbers.PushBack(i)
	}

	e := numbers.Front()
	for e != nil {
		tmp := e
		e = e.Next()
		if tmp.Value.(int)%2 == 0 {
			numbers.Remove(tmp)
		}
	}

	for e := numbers.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
```

## Complexity

Adding new items to the front or back of the linked list has the time complexity of O(1).

Deleting the first item is also O(1). Deleting the last item in a singly linked list is O(n), because the node before the last node must be found, and for that, every node must be visited. Deleting the last item can be done in O(1) time in a doubly linked list since nodes are connected to the previous and next nodes, so we can simply find the node before the last node and remove its reference to the next node.

## Application

Linked lists can be useful where the order of items matters, especially if there is a need for flexible reordering of the items or having current, next, and previous items. For example music players are a good candidate to implement using linked lists because they have playlists that play tracks in a predetermined order. At any given time, one track is playing while changing the current track with the previous or next ones is possible.

## Rehearsal

* [Linked List Serialization](./serialization_test.go), [Solution](./serialization.go)
* [Reverse a Linked List In-place](./reverse_in_place_test.go), [Solution](./reverse_in_place.go)
* [Join Two Sorted Linked Lists](./join_sorted_lists_test.go), [Solution](./join_sorted_lists.go)
* [Keep Repetitions](./keep_repetitions_test.go), [Solution](./keep%20repetitions.go)
* [Copy Linked List with Random Pointer](./copy_linklist_with_random_pointer_test.go), [Solution](./copy_linklist_with_random_pointer.go)
* [LRU Cache](./lru_cache_test.go), [Solution](./lru_cache.go)
