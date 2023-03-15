package linkedlist

import (
	"testing"
)

func TestSerializeAndUnserializeLinkedList(t *testing.T) {
	var tests = []string{
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
