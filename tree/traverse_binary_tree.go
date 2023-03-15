package tree

// TraverseBinaryTree performs in-order, pre-order, and post-order traversals on
// an integer binary tree and adds each visited node into a slice of integers
func TraverseBinaryTree(root *BinaryTreeNode) ([]int, []int, []int) {
	var (
		in, pre, post = []int{}, []int{}, []int{}
		inOrder       func(node *BinaryTreeNode)
		preOrder      func(node *BinaryTreeNode)
		postOrder     func(node *BinaryTreeNode)
	)

	inOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		in = append(in, node.Val)
		inOrder(node.Right)
	}
	preOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		pre = append(pre, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	postOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		post = append(post, node.Val)
	}

	inOrder(root)
	preOrder(root)
	postOrder(root)

	return in, pre, post
}
