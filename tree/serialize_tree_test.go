package tree

import "testing"

/*
TestSerializeAndUnserializeBinaryTree tests solution(s) with the following signature and problem description:

			type BinaryTreeNode struct {
				Val int
				Left *BinaryTreeNode
				Right *BinaryTreeNode
			}

			func Serialize(root *BinaryTreeNode) string
			func Deserialize(s string) *BinaryTreeNode


	     4
	    / \
	   /   \
	  2     6
	 / \   / \
	    3 5

Given a binary tree, write a serialize and deserialize function that turns it into a string representation
and back.

For example given `4,2,6,nil,3,5,nil` representing the above tree, deserialize it to a *BinaryTreeNode,
and given the *BinaryTreeNode serialize it back to `4,2,6,nil,3,5,nil`.
*/
func TestSerializeAndUnserializeBinaryTree(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1,2,3",
		"1,2,nil,4",
		"1,2,3,4,nil,5,6",
		"1,2,nil,4,nil,5,6",
	}
	for i, test := range tests {
		if got := Serialize(Deserialize(test)); got != test {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test, got)
		}
	}
}
