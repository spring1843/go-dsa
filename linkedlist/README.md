# Linked List

Linked lists are a collection of nodes, each of which is capable of storing at least one data element and is linked to the next node via a reference. One of the key advantages of linked lists over [arrays](../array) is their dynamic size, which allows for items to be added or removed without necessitating the resizing or shifting of other elements.

Two types of linked lists exist: singly linked lists, in which each node is linked only to the next node, and doubly linked lists, which enable each node to be connected to both the next and previous nodes.

## Implementation

When implementing linked lists, you typically have two concepts, the linked list itself, which has a link to the head and bottom nodes, and node(s), which contains at least one piece of data that is being stored in the linked list and also a link to the next node for singly linked lists, and an additional link to the previous node in doubly linked lists. To have a linked list in Go, you can define a node `struct` like the one below and use a pointer to the same structure. The example below creates a singly linked list that looks like `1->2->3->4->5`

When implementing linked lists, two key concepts come into play: the linked list itself, which includes references to the head and tail nodes, and the individual nodes which hold:

* At least one piece of data, such as name or number
* A reference to the next node
* A reference to the previous node in doubly linked lists

To create a linked list in Go, a node structure can be defined, and a pointer to the same structure can be used. For instance, a singly linked list that follows the format 1->2->3->4->5 can be created using the following example:

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

In a singly linked list, a pointer is typically stored for the first and last items. Adding items to the front or back of the list is a constant-time operation. However, deleting the last item can be challenging, as the last item's pointer needs to be updated to the second-to-last item. This is where having a reference to the last item in each node proves useful. In contrast, doubly linked lists maintain pointers to both the previous and next nodes, which makes deletion operations less expensive.

The Go standard library contains an implementation of [doubly linked lists](https://golang.org/pkg/container/list/). In the following example, numbers from 1 to 10 are added to the list, and then even numbers are removed, and the resulting linked list containing odd numbers is printed.

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

Deleting the first item is also O(1). Deleting the last item in a singly linked list is O(n), because we have to find the node before the last node, and for that, we have to visit every node. In a doubly linked list where a node is connected to both previous and next nodes, though, deleting the last item can be done in O(1) time.

## Application

Linked lists can be useful where the order of similar items matters, especially if there is a need for flexible reordering of the items or having current, next, and previous items.

A music album is a good example of linked lists. The tracks come in an order, and at any given time, you are listening to one track while you have the option of going to the next song in a singly linked list and both the previous and next song in a doubly linked list.

## Rehearsal

* [Linked List Serialization](./serialization_test.go), [Solution](./serialization.go)
* [Reverse a Linked List In-place](./reverse_in_place_test.go), [Solution](./reverse_in_place.go)
* [Join Two Sorted Linked Lists](./join_sorted_lists_test.go), [Solution](./join_sorted_lists.go)
* [Keep Repetitions](./keep_repetitions_test.go), [Solution](./keep_repetitions_test.go)
* [Copy Linked List with Random Pointer](./copy_linklist_with_random_pointer_test.go), [Solution](./copy_linklist_with_random_pointer.go)
* [Implement LRU Cache](./lru_cache_test.go), [Solution](./lru_cache.go)
