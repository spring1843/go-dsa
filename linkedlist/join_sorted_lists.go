package linkedlist

// JoinTwoSortedLinkedLists merges two sorted linked lists into one.
func JoinTwoSortedLinkedLists(l1, l2 *Node) *Node {
	head := &Node{}
	last := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			last, l1 = moveTolast(last, l1)
		} else {
			last, l2 = moveTolast(last, l2)
		}
	}

	for l1 != nil {
		last, l1 = moveTolast(last, l1)
	}

	for l2 != nil {
		last, l2 = moveTolast(last, l2)
	}

	return head.Next
}

func moveTolast(last, node *Node) (*Node, *Node) {
	last.Next = node
	last = last.Next
	node = node.Next
	return last, node
}
