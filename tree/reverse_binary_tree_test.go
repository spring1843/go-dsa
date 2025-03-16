package tree

import "testing"

/*
TestReverseBinaryTree tests solution(s) with the following signature and problem description:

	func ReverseBinaryTree(root *BinaryTreeNode) *BinaryTreeNode

Given a binary tree, reverse it.

	     1
	   /   \
	  2     3
	 / \   / \
	4   5 6   7

For example given the above tree, return:

	     1
	   /   \
	  3     2
	 / \   / \
	7   6 5   4

Each tree is represented as a string by "1,2,3,4,5,6,7" and "1,3,2,7,6,5,4" respectively.
*/
func TestReverseBinaryTree(t *testing.T) {
	tests := []struct {
		tree, reversed string
	}{
		{"", ""},
		{"1,2,3", "1,3,2"},
		{"1,2,nil", "1,nil,2"},
		{"1,2,3,4,5,6,7", "1,3,2,7,6,5,4"},
		{"1,2,3,4", "1,3,2,nil,nil,nil,4"},
	}

	for i, test := range tests {
		if got := Serialize(ReverseBinaryTree(Deserialize(test.tree))); got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.reversed, got)
		}
	}
}
