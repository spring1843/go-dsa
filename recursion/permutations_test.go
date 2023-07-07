package recursion

import (
	"reflect"
	"testing"
)

/*
TestPermutations tests solution(s) with the following signature and problem description:

	func Permutations(numbers []int) [][]int

Returns all possible permutations of the given integers.
*/
func TestPermutations(t *testing.T) {
	tests := []struct {
		input        []int
		permutations [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
		{[]int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 3, 2}, {1, 4, 2, 3}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 3, 2, 1}, {4, 3, 1, 2}, {4, 1, 3, 2}, {4, 1, 2, 3}}},
	}

	for i, test := range tests {
		if got := Permutations(test.input); !reflect.DeepEqual(got, test.permutations) {
			t.Fatalf("Failed test case %#v. Want %d got %#v", i, test.permutations, got)
		}
	}
}
