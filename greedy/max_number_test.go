package greedy

import (
	"reflect"
	"testing"
)

/*
TestMaxNumber tests solution(s) with the following signature and problem description:

	func MaxNumber(number1, number2 []int, n int) []int

Given two numbers represented as list of positive integers, and an integer n, return the
largest possible integer with n digits that can be constructed by merging digits from two
numbers while respecting the order of digits in each number.

For example given {5, 4, 3, 2, 1}, {9, 8, 7, 6} and 9, it should return a 9 digit number
{9, 8, 7, 6, 5, 4, 3, 2, 1}.
*/
func TestMaxNumber(t *testing.T) {
	tests := []struct {
		number1       []int
		number2       []int
		k             int
		largestNumber []int
	}{
		{[]int{}, []int{}, 1, []int{}},
		{[]int{}, []int{1}, 1, []int{1}},
		{[]int{0}, []int{1}, 1, []int{1}},
		{[]int{0}, []int{1}, 2, []int{1, 0}},
		{[]int{1}, []int{0}, 2, []int{1, 0}},
		{[]int{1}, []int{2}, 4, []int{}},
		{[]int{2}, []int{1}, 4, []int{}},
		{[]int{3, 2, 1}, []int{1}, 3, []int{3, 2, 1}},
		{[]int{6, 9}, []int{6, 0, 4}, 5, []int{6, 9, 6, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 4, []int{3, 1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{9, 8, 7, 6}, 9, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{9, 8, 7, 6}, 8, []int{9, 8, 7, 6, 5, 4, 3, 2}},
	}

	for i, test := range tests {
		if got := MaxNumber(test.number1, test.number2, test.k); !reflect.DeepEqual(got, test.largestNumber) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.largestNumber, got)
		}
	}
}
