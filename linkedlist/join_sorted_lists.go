package linkedlist

// JoinTwoSortedLinkedLists solves the problem in O(n) time and O(1) space.
func JoinTwoSortedLinkedLists(l1, l2 *Node) *Node {
	head := &Node{Val: 0, Next: nil}
	last := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			last, l1 = moveToLast(last, l1)
		} else {
			last, l2 = moveToLast(last, l2)
		}
	}

	for l1 != nil {
		last, l1 = moveToLast(last, l1)
	}

	for l2 != nil {
		last, l2 = moveToLast(last, l2)
	}

	return head.Next
}

func moveToLast(last, node *Node) (*Node, *Node) {
	last.Next = node
	last = last.Next
	node = node.Next
	return last, node
}
