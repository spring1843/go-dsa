package linkedlist

import (
	"math"
)

// KeepRepetitions solves the problem in O(n) time and O(1) space.
func KeepRepetitions(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	var output, tmp *Node
	lastValue := math.MinInt
	for head != nil {
		if lastValue == head.Val {
			if tmp == nil {
				tmp = NewNode(head.Val)
				output = tmp
				continue
			}
			if tmp.Val != head.Val {
				tmp.Next = NewNode(head.Val)
				tmp = tmp.Next
			}
		}
		lastValue = head.Val
		head = head.Next
	}

	return output
}
