package tree

import (
	"strconv"
	"strings"
)

const (
	delimiter = ","
	nilNode   = "nil"
)

// BinaryTreeNode is a node in a Binary Tree of integers
type BinaryTreeNode struct {
	// Val is an integer value of a node in a Binary Tree
	Val int

	// Left is the left node in a binary tree
	Left *BinaryTreeNode

	// Right is the right node in a binary tree
	Right *BinaryTreeNode
}

// Serialize serializes a given binary tree into a string
func Serialize(root *BinaryTreeNode) string {
	if root == nil {
		return ""
	}

	result := ""
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = result + nilNode + delimiter
		} else {
			result = result + strconv.Itoa(node.Val) + delimiter
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
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

// Unserialize unserializes a given string into a binary tree
func Unserialize(s string) *BinaryTreeNode {
	if s == "" {
		return nil
	}

	nodes := []*BinaryTreeNode{}

	for _, node := range strings.Split(s, delimiter) {
		if node == nilNode {
			nodes = append(nodes, nil)
		} else {
			Value, _ := strconv.Atoi(node)
			nodes = append(nodes, &BinaryTreeNode{Val: Value})
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
			node.Left = nodes[i]
		}
		i++

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.Right = nodes[i]
		}
		i++
	}

	return nodes[0]
}
