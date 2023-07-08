package tree

import (
	"strings"
	"testing"
)

/*
TestEvaluateBinaryExpressionTree tests solution(s) with the following signature and problem description:

	func EvaluateBinaryExpressionTree(node *stringBinaryTreeNode) (float64, error)

Given an expression binary tree like (D) in _Figure2_ evaluate it to a float64 like 100 allowing four
arithmetic operations.
*/
func TestEvaluateBinaryExpressionTree(t *testing.T) {
	tests := []struct {
		tree        string
		result      float64
		expectedErr bool
	}{
		{"", 0.0, false},
		{"*,6,2", 12.0, false},
		{"/,6,2", 3.0, false},
		{"-,6,2", 4.0, false},
		{"+,*,6,+,2,nil,nil,3,4", 20.0, false},
		{"*,*,+,5,5,2,2", 100.0, false}, // (D) in Figure 2
		{"*,A,2", -1, true},
		{"*,2,B", -1, true},
		{"A,1,2", -1, true},
	}

	for i, test := range tests {
		got, err := EvaluateBinaryExpressionTree(unserializeStringBinaryTree(test.tree))
		if !test.expectedErr && err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if test.expectedErr && err == nil {
			t.Fatal("did not get expected error")
		}
		if got != test.result {
			t.Fatalf("Failed test case #%d. Want %#f got %#f", i, test.result, got)
		}
	}
}

func TestSerializeAndUnserializeStringBinaryTree(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1,2,3",
		"1,2,nil,4",
		"1,2,3,4,nil,5,6",
		"1,2,nil,4,nil,5,6",
		"a,b,nil,c,nil,d,e",
		"+,b,nil,a,nil,c,e",
	}
	for i, test := range tests {
		if got := serializeStringBinaryTree(unserializeStringBinaryTree(test)); got != test {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test, got)
		}
	}
}

func serializeStringBinaryTree(root *stringBinaryTreeNode) string {
	if root == nil {
		return ""
	}

	result := ""
	queue := []*stringBinaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = result + nilNode + delimiter
		} else {
			result = result + node.val + delimiter
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}
	}

	for strings.HasSuffix(result, nilNode+delimiter) {
		result = strings.TrimSuffix(result, nilNode+delimiter)
	}
	for strings.HasSuffix(result, delimiter) {
		result = strings.TrimSuffix(result, delimiter)
	}

	return result
}

func unserializeStringBinaryTree(s string) *stringBinaryTreeNode {
	if s == "" {
		return nil
	}

	nodes := []*stringBinaryTreeNode{}

	for _, node := range strings.Split(s, delimiter) {
		if node == nilNode {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &stringBinaryTreeNode{val: node})
		}
	}

	i := 1
	for _, node := range nodes {
		if node == nil {
			continue
		}

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.left = nodes[i]
		}
		i++

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.right = nodes[i]
		}
		i++
	}

	return nodes[0]
}
