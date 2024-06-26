package linkedlist

import "testing"

/*
TestSerializeAndUnserializeLinkedList tests solution(s) with the following signature and problem description:

	func Serialize(node *Node) string
	func Unserialize(stringRepresentation string) *Node

Write a function that turns a linked list into a string representation, and then a function that turns that
string representation to an actual linked list.
*/
func TestSerializeAndUnserializeLinkedList(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1->2",
		"1->2->3->4->2->1",
	}
	for i, test := range tests {
		got := Serialize(Unserialize(test))
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
