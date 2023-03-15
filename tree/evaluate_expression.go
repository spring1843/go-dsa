package tree

import "strconv"

type stringBinaryTreeNode struct {
	val   string
	left  *stringBinaryTreeNode
	right *stringBinaryTreeNode
}

// EvaluateBinaryExpressionTree evaluates a binary expression tree and evaluates
// it allowing the four basic arithmetic operations.
func EvaluateBinaryExpressionTree(node *stringBinaryTreeNode) float64 {
	if node == nil || node.val == "" {
		return 0
	}
	switch node.val {
	case "*":
		return EvaluateBinaryExpressionTree(node.left) * EvaluateBinaryExpressionTree(node.right)
	case "/":
		return EvaluateBinaryExpressionTree(node.left) / EvaluateBinaryExpressionTree(node.right)
	case "+":
		return EvaluateBinaryExpressionTree(node.left) + EvaluateBinaryExpressionTree(node.right)
	case "-":
		return EvaluateBinaryExpressionTree(node.left) - EvaluateBinaryExpressionTree(node.right)
	}

	val, err := strconv.ParseFloat(node.val, 64)
	if err != nil {
		panic(err)
	}
	return val
}
