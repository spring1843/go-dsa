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

Write a function that turns a linked list into a string representation (Serialize), and then a function
that turns that string representation to an actual linked list (Deserialize).
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
