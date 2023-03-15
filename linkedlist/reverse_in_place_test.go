package linkedlist

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	var tests = []struct {
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
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.reversed, got)
		}
	}
}
