package tree

// BalancedBinarySearchTree solves the problem in O(n) time and O(n) space.
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
