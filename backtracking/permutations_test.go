package backtracking

import (
	"reflect"
	"testing"
)

/*
TestPermutations tests solution(s) with the following signature and problem description:

	func Permutations(input []int) [][]int

Given a list of integers like {1,2}, produce all possible combinations like {1,2},{2,1}.
*/
func TestPermutations(t *testing.T) {
	tests := []struct {
		nums         []int
		permutations [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
		{[]int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 3, 2}, {1, 4, 2, 3}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 3, 2, 1}, {4, 3, 1, 2}, {4, 1, 3, 2}, {4, 1, 2, 3}}},
	}

	for i, test := range tests {
		if got := Permutations(test.nums); !reflect.DeepEqual(test.permutations, got) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.permutations, got)
		}
	}
}
