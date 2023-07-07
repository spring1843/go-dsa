package array

import (
	"reflect"
	"testing"
)

/*
TestAddTwoNumbers tests solution(s) with the following signature and problem description:

	AddTwoNumbers(num1, num2 []int) []int

Adds two numbers which are represented as an array and returns the results
In the same format. For example [2,5], and [3,5] would add up to[6,0] because 25 + 35 = 60.
*/
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		num1, num2, sum []int
	}{
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{1}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1}, []int{9}, []int{1, 0}},
		{[]int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}},
	}
	for i, test := range tests {
		if got := AddTwoNumbers(test.num1, test.num2); !reflect.DeepEqual(got, test.sum) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sum, got)
		}
	}
}
