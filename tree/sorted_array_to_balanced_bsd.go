package tree

// BalancedBinarySearchTree takes a sorted list of numbers and returns a balanced binary search tree
func BalancedBinarySearchTree(sorted []int) *BinaryTreeNode {
	if len(sorted) == 0 {
		return nil
	}

	mid := len(sorted) / 2
	root := &BinaryTreeNode{Val: sorted[mid]}
	root.Left = BalancedBinarySearchTree(sorted[:mid])
	root.Right = BalancedBinarySearchTree(sorted[mid+1:])

	return root
}
