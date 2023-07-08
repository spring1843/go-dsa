package linkedlist

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

const (
	randomValueSeparator = ":"
	noRandomLink         = "nil"
)

/*
TestCopyLinkedListWithRandomPointer tests solution(s) with the following signature and problem description:

	func CopyLinkedListWithRandomPointer(head *RandomNode) *RandomNode

Given a singly connected linked list in which each node may optionally be connected to another node in
random order, return a deep copy of the linked list.
*/
func TestCopyLinkedListWithRandomPointer(t *testing.T) {
	tests := []struct {
		list string
	}{
		{""},
		{"1:nil"}, // What follows after : is the index of the node the random link is connected to
		{"1:1"},
		{"1:nil->2:1->3:4->4:nil"},
		{"1:nil->2:1->3:4->4:nil"},
		{"1:nil->2:1->2:nil->3:4->4:nil"},
	}

	for i, test := range tests {
		list := unserializeRandomNode(test.list)
		deepCopy := CopyLinkedListWithRandomPointer(list)

		if list != nil {
			list.Random = &RandomNode{Val: math.MinInt64, Next: nil}
			if deepCopy.Random != nil && deepCopy.Random.Val == math.MinInt64 {
				t.Fatal("Not a deep copy. Changing the value in the original list, changed the value in whats expected to be a deep copy")
			}
		}

		got := serializeRandomNode(deepCopy)
		if got != test.list {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.list, got)
		}
	}
}

func serializeRandomNode(node *RandomNode) string {
	if node == nil {
		return ""
	}

	indices := indicesMap(node)

	outputNode := func(n *RandomNode) string {
		output := strconv.Itoa(n.Val)
		if n.Random == nil {
			output += randomValueSeparator + noRandomLink
		} else {
			output += randomValueSeparator + strconv.Itoa(indices[n.Random])
		}
		output += separator
		return output
	}

	output := outputNode(node)
	for node.Next != nil {
		node = node.Next
		output += outputNode(node)
	}

	return strings.TrimSuffix(output, separator)
}

func indicesMap(node *RandomNode) map[*RandomNode]int {
	indices := make(map[*RandomNode]int)
	i := 1
	indices[node] = i
	for node.Next != nil {
		i++
		node = node.Next
		indices[node] = i
	}
	return indices
}

func unserializeRandomNode(stringRepresentation string) *RandomNode {
	if stringRepresentation == "" {
		return nil
	}
	var cur, last *RandomNode
	broken := strings.Split(stringRepresentation, separator)

	nodes := make(map[int]*RandomNode)
	for i := len(broken) - 1; i >= 0; i-- {
		last = cur
		val := broken[i][0:strings.Index(broken[i], randomValueSeparator)]
		cur = &RandomNode{
			Val:  atoi(val),
			Next: nil,
		}
		if last != nil {
			cur.Next = last
		}
		nodes[i] = cur
	}

	head := cur

	for i := 0; i < len(broken); i++ {
		randomIndex := broken[i][strings.Index(broken[i], randomValueSeparator)+1:]
		if randomIndex != noRandomLink {
			cur.Random = nodes[atoi(randomIndex)-1]
		}
		cur = cur.Next
	}

	return head
}
