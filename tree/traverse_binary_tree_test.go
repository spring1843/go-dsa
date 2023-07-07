package tree

import (
	"reflect"
	"testing"
)

/*
TestTraverseBinaryTree tests solution(s) with the following signature and problem description:

	func TraverseBinaryTree(root *BinaryTreeNode) ([]int, []int, []int)

Performs in-order, pre-order, and post-order traversals on an integer binary tree and adds
each visited node into a slice of integers.
*/
func TestTraverseBinaryTree(t *testing.T) {
	tests := []struct {
		tree          string
		in, pre, post []int
	}{
		{"", []int{}, []int{}, []int{}},
		{"1", []int{1}, []int{1}, []int{1}},
		{"4,2,6,1,3,5,7", []int{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 3, 2, 5, 7, 6, 4}},
		{"3,1,2", []int{1, 3, 2}, []int{3, 1, 2}, []int{1, 2, 3}},
		{"4,2,6,nil,3,5", []int{2, 3, 4, 5, 6}, []int{4, 2, 3, 6, 5}, []int{3, 2, 5, 6, 4}},
	}
	for i, test := range tests {
		gotIn, gotPre, gotPost := TraverseBinaryTree(Unserialize(test.tree))
		if !reflect.DeepEqual(gotIn, test.in) {
			t.Fatalf("Failed in-order test case #%d.  Want %#v got %#v", i, test.in, gotIn)
		}
		if !reflect.DeepEqual(gotPre, test.pre) {
			t.Fatalf("Failed pre-order test case #%d. Want %#v got %#v", i, test.pre, gotPre)
		}
		if !reflect.DeepEqual(gotPost, test.post) {
			t.Fatalf("Failed post-order test case #%d. Want %#v got %#v", i, test.post, gotPost)
		}
	}
}
