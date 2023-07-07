package queue

import (
	"reflect"
	"testing"
)

/*
TestMaxOfKLengthSubArrays tests solution(s) with the following signature and problem description:

	func MaxOfKLengthSubArrays(numbers []int, k int) ([]int, error) {

takes a list of integers and returns a list of the maximum of each sub-array of length k.
*/
func TestMaxOfKLengthSubArrays(t *testing.T) {
	tests := []struct {
		k        int
		list     []int
		maximums []int
	}{
		{3, []int{5, 3, 5, 6, 7, 8}, []int{5, 6, 7, 8}},
		{2, []int{5, 3, 5, 6, 7, 8}, []int{5, 5, 6, 7, 8}},
		{3, []int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	}

	for i, test := range tests {
		got, err := MaxOfKLengthSubArrays(test.list, test.k)
		if err != nil {
			t.Fatalf("Unexpected error occurred. Error : %s", err)
		}
		if !reflect.DeepEqual(got, test.maximums) {
			t.Fatalf("Failed test #%d, Failed getting list of maximums. Want %#v got %#v", i, test.maximums, got)
		}
	}
}
