package dp

import (
	"testing"
)

/*
TestSumUpToInteger tests solution(s) with the following signature and problem description:

	func SumUpToInteger(numbers []int, sum int) bool

Given a set of positive integers like {1,2,3,4,5} and an integer like 7 write a
function that returns true if there are two numbers in the list that sum up to the given
integer and false otherwise.
*/
func TestSumUpToInteger(t *testing.T) {
	tests := []struct {
		numbers []int
		sum     int
		sumsUp  bool
	}{
		{[]int{1}, 1, true},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2}, 3, true},
		{[]int{1, 2}, 4, false},
		{[]int{1, 2, 3, 4, 5}, 7, true},
		{[]int{1, 2, 3, 4, 5}, 8, true},
		{[]int{1, 2, 3, 4, 5}, 15, true},
		{[]int{1, 2, 3, 4, 5}, 16, false},
		{[]int{1, 5, 8, 9, 10, 20, 30}, 25, true},
	}

	for i, test := range tests {
		if got := SumUpToInteger(test.numbers, test.sum); got != test.sumsUp {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.sumsUp, got)
		}
	}
}
