package queue

import (
	"testing"

	"github.com/spring1843/go-dsa/tree"
)

/*
TestIsTreeSymmetrical tests solution(s) with the following signature and problem description:

	func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error)

Given a binary tree the one above "2,4,4,5,6,6,5", true if it is symmetric (mirrored from the root), and false otherwise.

		     2
		   /   \
		  /     \
		 4       4
		/ \     / \
	   5   6   6   5
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
		got, err := IsTreeSymmetrical(tree.Unserialize(test.tree))
		if err != nil {
			t.Fatalf("Failed test case #%d. Unexpected error %s", i, err)
		}
		if got != test.isSymmetric {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.isSymmetric, got)
		}
	}
}
