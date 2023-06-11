package tree

import (
	"strings"
	"testing"
)

func TestEvaluateBinaryExpressionTree(t *testing.T) {
	tests := []struct {
		tree   string
		result float64
	}{
		{"", 0.0},
		{"*,6,2", 12.0},
		{"/,6,2", 3.0},
		{"-,6,2", 4.0},
		{"+,*,6,+,2,nil,nil,3,4", 20.0},
		{"*,*,+,5,5,2,2", 100.0}, // (D) in Figure 2
	}

	for i, test := range tests {
		if got := EvaluateBinaryExpressionTree(unserializeStringBinaryTree(test.tree)); got != test.result {
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
