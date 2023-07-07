package heap

import (
	"testing"

	"github.com/spring1843/go-dsa/linkedlist"
)

/*
TestMergeSortedLists tests solution(s) with the following signature and problem description:

	func MergeSortedLists(lists []*linkedlist.Node) *linkedlist.Node

Merges multiple sorted singly linked lists into one.
*/
func TestMergeSortedLists(t *testing.T) {
	tests := []struct {
		sortedLinkedLists []string
		merged            string
	}{
		{[]string{""}, ""},
		{[]string{"1"}, "1"},
		{[]string{"1", ""}, "1"},
		{[]string{"1", "1->2"}, "1->1->2"},
		{[]string{"1", "1->2", "3->4->5"}, "1->1->2->3->4->5"},
		{[]string{"1", "1->2", "1->2->3", "1->3->5", "1->2->4"}, "1->1->1->1->1->2->2->2->3->3->4->5"},
	}

	for i, test := range tests {
		nodes := []*linkedlist.Node{}
		for _, sortedLinkedList := range test.sortedLinkedLists {
			nodes = append(nodes, linkedlist.Unserialize(sortedLinkedList))
		}
		got := linkedlist.Serialize(MergeSortedLists(nodes))
		if got != test.merged {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.merged, got)
		}
	}
}
