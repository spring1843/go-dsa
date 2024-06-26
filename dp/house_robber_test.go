package dp

import "testing"

/*
TestMaxHouseRobber tests solution(s) with the following signature and problem description:

	func MaxHouseRobber(wealth []int) int

Given an array representing the amount of wealth inside houses like {1,2,3,4} return the
maximum wealth the robber can steal like `6` with the condition that he is not allowed to
rob two consecutive houses.
*/
func TestMaxHouseRobber(t *testing.T) {
	tests := []struct {
		houses []int
		max    int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3, 4}, 6},
		{[]int{1, 3, 5, 10, 15}, 21},
		{[]int{1, 5, 8, 9, 10}, 19},
		{[]int{1, 5, 8, 9, 10, 20, 30}, 49},
	}

	for i, test := range tests {
		if got := MaxHouseRobber(test.houses); got != test.max {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.max, got)
		}
	}
}
