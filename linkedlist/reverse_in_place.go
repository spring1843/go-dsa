package linkedlist

// ReverseLinkedList solves the problem in O(n) time and O(1) space.
func ReverseLinkedList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var last *Node
	for head != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}

	return last
}
