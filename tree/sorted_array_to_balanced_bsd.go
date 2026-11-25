package tree

// BalancedBinarySearchTree solves the problem in O(n) time and O(n) space.
func BalancedBinarySearchTree(sorted []int) *BinaryTreeNode {
	if len(sorted) == 0 {
		return nil
	}

	mid := len(sorted) / 2

	return &BinaryTreeNode{
		Val:   sorted[mid],
		Left:  BalancedBinarySearchTree(sorted[:mid]),
		Right: BalancedBinarySearchTree(sorted[mid+1:]),
	}
}
