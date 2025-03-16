package queue

import (
	"testing"

	"github.com/spring1843/go-dsa/tree"
)

/*
TestIsTreeSymmetrical tests solution(s) with the following signature and problem description:

	func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error)

Given a binary tree return true of it is symmetric and false otherwise. A tree is symmetric if you
can draw a vertical line through the root and then the left subtree is the mirror image of the right subtree.

	  Symmetric       Not Symmetric
	      2                2
	    /   \            /   \
	   /     \          /     \
	  4       4        3       4
	 / \     / \      / \     / \
	5   6   6   5    5   6   6   5

For example given "2,4,4,5,6,6,5", shown in the symmetric tree above return true.
Given "2,3,4,5,6,6,5", shown in the not symmetric tree above return false.
*/
func TestIsTreeSymmetrical(t *testing.T) {
	tests := []struct {
		tree        string
		isSymmetric bool
	}{
		{"", false},
		{"1", true},
		{"1,2,2", true},
		{"1,2,3", false},
		{"1,2,2,3,nil,nil,3", true},
		{"1,2,nil,4", false},
		{"1,2,3,4,nil,5,6", false},
		{"1,2,nil,4,nil,5,6", false},
		{"2,4,4,5,6,5,6", false},
		{"2,4,4,5,6,6,5", true},
	}
	for i, test := range tests {
		got, err := IsTreeSymmetrical(tree.Deserialize(test.tree))
		if err != nil {
			t.Fatalf("Failed test case #%d. Unexpected error %s", i, err)
		}
		if got != test.isSymmetric {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.isSymmetric, got)
		}
	}
}
