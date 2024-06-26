package dnc

import "testing"

/*
TestBinarySearch tests solution(s) with the following signature and problem description:

	func BinarySearch(list []int, search int) int

Given a sorted set of integers like {1,2,3,4,6}, and a target int like 4 find its
position in the set like 3 using Binary Search.

In Binary Search we start with the middle element of the set and compare it with
the target. If the target is greater than the middle element we search the right
half of the set, otherwise we search the left half of the set. We repeat this
process until we find the target or we exhaust the set.
*/
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		sortedNumbers    []int
		search, position int
	}{
		{[]int{}, 3, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 5, 10}, 4, -1},
		{[]int{1, 2, 3, 8, 10}, 8, 3},
		{[]int{1, 2, 3, 8, 10, 11}, 8, 3},
		{[]int{1, 2, 3, 4, 5, 9}, 4, 3},
	}

	for i, test := range tests {
		if got := BinarySearch(test.sortedNumbers, test.search); got != test.position {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.position, got)
		}
	}
}
