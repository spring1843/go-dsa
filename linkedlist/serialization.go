package linkedlist

import (
	"strconv"
	"strings"
)

const separator = `->`

// Node is a link in a singly linked list that stores integers.
type Node struct {
	// Val is the value of the node
	Val int

	// Next is a link to the next node
	Next *Node
}

// NewNode returns a new node.
func NewNode(v int) *Node {
	return &Node{
		Val:  v,
		Next: nil,
	}
}

// Serialize solves the problem in O(n) time and O(1) space.
func Serialize(node *Node) string {
	if node == nil {
		return ""
	}

	output := strconv.Itoa(node.Val) + separator
	for node.Next != nil {
		node = node.Next
		output += strconv.Itoa(node.Val) + separator
	}

	return strings.TrimSuffix(output, separator)
}

// Deserialize solves the problem in O(n) time and O(1) space.
func Deserialize(stringRepresentation string) *Node {
	if stringRepresentation == "" {
		return nil
	}
	var cur, last *Node
	broken := strings.Split(stringRepresentation, separator)

	for i := len(broken) - 1; i >= 0; i-- {
		last = cur
		cur = NewNode(atoi(broken[i]))
		if last != nil {
			cur.Next = last
		}
	}

	return cur
}

func atoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		return -1
	}
	return i
}
