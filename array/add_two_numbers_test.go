package array

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
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
