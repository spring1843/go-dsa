package linkedlist

import (
	"testing"
)

/*
TestReverseLinkedList tests solution(s) with the following signature and problem description:

	func ReverseLinkedList(head *Node) *Node

Reverse a given linked list in place.
*/
func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		list, reversed string
	}{
		{"", ""},
		{"1", "1"},
		{"1->2", "2->1"},
		{"1->2->3", "3->2->1"},
	}

	for i, test := range tests {
		got := Serialize(ReverseLinkedList(Unserialize(test.list)))
		if got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.reversed, got)
		}
	}
}
