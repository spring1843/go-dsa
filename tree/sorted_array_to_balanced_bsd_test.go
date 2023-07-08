package tree

import "testing"

/*
TestBalancedBinarySearchTree tests solution(s) with the following signature and problem description:

	func BalancedBinarySearchTree(sorted []int) *BinaryTreeNode

Given a sorted array of integers like {1,2,3,4,5} return the root to a balanced BST.
*/
func TestBalancedBinarySearchTree(t *testing.T) {
	tests := []struct {
		numbers []int
		tree    string
	}{
		{[]int{}, ""},
		{[]int{1}, "1"},
		{[]int{1, 2}, "2,1"},
		{[]int{1, 2, 3}, "2,1,3"},
		{[]int{1, 2, 3, 4}, "3,2,4,1"},
		{[]int{1, 2, 3, 4, 5}, "3,2,5,1,nil,4"},
		{[]int{1, 2, 3, 4, 6}, "3,2,6,1,nil,4"},
		{[]int{1, 2, 3, 4, 6, 7}, "4,2,7,1,3,6"},
		{[]int{1, 2, 3, 4, 6, 7, 8}, "4,2,7,1,3,6,8"},
	}

	for i, test := range tests {
		if got := Serialize(BalancedBinarySearchTree(test.numbers)); got != test.tree {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.tree, got)
		}
	}
}
