package tree

import "testing"

/*
TestSerializeAndUnserializeBinaryTree tests solution(s) with the following signature and problem description:

	func Unserialize(s string) *BinaryTreeNode

Unserializes a given string into a binary tree.
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
