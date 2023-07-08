package tree

// ReverseBinaryTree solves the problem in O(n) time and O(1) space.
func ReverseBinaryTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = ReverseBinaryTree(root.Right), ReverseBinaryTree(root.Left)
	return root
}
