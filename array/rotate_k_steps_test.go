package array

import (
	"slices"
	"testing"
)

/*
TestRotateKSteps tests solution(s) with the following signature and problem description:

	RotateKSteps(list []int, k int)

Given an list of integers and a number k, rotate the array k times.

For example given {1,2,3} and 3, return {1,2,3} because:

After 1 rotation: {3,1,2}.
After 2 rotations: {2,3,1}.
After 3 rotations: {1,2,3}.
*/
func TestRotateKSteps(t *testing.T) {
	tests := []struct {
		list        []int
		steps       int
		rotatedList []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
	}

	for i, test := range tests {
		RotateKSteps(test.list, test.steps)
		if !slices.Equal(test.list, test.rotatedList) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.rotatedList, test.list)
		}
	}
}
