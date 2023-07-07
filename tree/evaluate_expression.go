package tree

import (
	"fmt"
	"strconv"
)

type stringBinaryTreeNode struct {
	val   string
	left  *stringBinaryTreeNode
	right *stringBinaryTreeNode
}

// EvaluateBinaryExpressionTree solves the problem in O(n) time and O(n) space.
func EvaluateBinaryExpressionTree(node *stringBinaryTreeNode) (float64, error) {
	if node == nil || node.val == "" {
		return 0, nil
	}

	left, err := EvaluateBinaryExpressionTree(node.left)
	if err != nil {
		return -1, fmt.Errorf("failed evaluating value of left node to %s. %s", node.val, err)
	}

	right, err := EvaluateBinaryExpressionTree(node.right)
	if err != nil {
		return -1, fmt.Errorf("failed evaluating value of right node to %s. %s", node.val, err)
	}

	switch node.val {
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	}

	val, err := strconv.ParseFloat(node.val, 64)
	if err != nil {
		return -1, fmt.Errorf("failed parsing value %s as float64. %s", node.val, err)
	}
	return val, nil
}
