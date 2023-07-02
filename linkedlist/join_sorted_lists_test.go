package linkedlist

import (
	"testing"
)

func TestJoinTwoSortedLinkedLists(t *testing.T) {
	tests := []struct {
		list1, list2, joined string
	}{
		{"", "", ""},
		{"1", "", "1"},
		{"", "1", "1"},
		{"1", "2", "1->2"},
		{"1", "2->3", "1->2->3"},
		{"1->3", "2", "1->2->3"},
		{"1->4->6", "2->3->5->7", "1->2->3->4->5->6->7"},
		{"1->4->5->6", "2->3->7", "1->2->3->4->5->6->7"},
		{"1->2", "1->2->3", "1->1->2->2->3"},
		{"1->4", "2->3->7", "1->2->3->4->7"},
	}

	for i, test := range tests {
		got := Serialize(JoinTwoSortedLinkedLists(Unserialize(test.list1), Unserialize(test.list2)))
		if got != test.joined {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.joined, got)
		}
	}
}
