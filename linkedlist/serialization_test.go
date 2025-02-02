package linkedlist

import "testing"

/*
TestSerializeAndDeserializeLinkedList tests solution(s) with the following signature and problem description:

	type Node struct {
		Val int
		Next *Node
	}

	func Serialize(node *Node) string
	func Deserialize(stringRepresentation string) *Node

Write a function that turns a linked list of integers into a string representation (Serialize) and  a function
that turns that string representation to an actual linked list (Deserialize).

For example consider the following example of a linked list containing numbers 1,2,3:

node1 := &Node{Val: 1}
node2 := &Node{Val: 2}
node3 := &Node{Val: 3}
node1.Next = node2
node2.Next = node3

A string representation of this linked-list should look like 1->2->3. This means that
the node with value 1 is connected to node with value 2, and node with value 2 is connected to node with value 3.
*/
func TestSerializeAndDeserializeLinkedList(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1->2",
		"1->2->3->4->2->1",
	}
	for i, test := range tests {
		got := Serialize(Deserialize(test))
		if got != test {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test, got)
		}
	}
}

func TestATOI(t *testing.T) {
	if got := atoi("A"); got != -1 {
		t.Fatalf("want %d, got %d", -1, got)
	}
}
