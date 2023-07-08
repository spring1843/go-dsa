package tree

import "testing"

/*
TestSerializeAndUnserializeBinaryTree tests solution(s) with the following signature and problem description:

			func Serialize(root *BinaryTreeNode) string
			func Unserialize(s string) *BinaryTreeNode


	     4
	    / \
	   /   \
	  2     6
	 / \   / \
	    3 5

Write two functions to serialize and unserialize a binary tree like the one above to and
from a string like `4,2,6,nil,3,5,nil`.
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
		if got := Serialize(Unserialize(test)); got != test {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test, got)
		}
	}
}
