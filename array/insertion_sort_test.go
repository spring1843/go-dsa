package array

import (
	"slices"
	"testing"
)

/*
TestInsertionSort tests solution(s) with the following signature and problem description:

	InsertionSort(input []int)

Given an array of unsorted integers, sort the array using the Insertion Sort algorithm.
The algorithm should be in-place, meaning it should not create a new array.

To use Insertion Sort to sort a slice increasingly:
1- Create a sorted list and an unsorted list. Put the first element of the list in the sorted and the rest of the elements in the unsorted list.
2- Compare the first element of the unsorted list with each element of the sorted list starting from the end and swapping as needed.
3- If there are remaining elements in the unsorted list keep doing 2.

The name Insertion Sort comes from the observation that each element is picked and inserted at the right place.
To save memory space in an in-place implementation of Insertion Sort we do not create two new lists, we rather use the left side and right side of the input list and an imaginary divisor as a virtual way of having two separate lists.

To sort {2,3,1} using Insertion Sort:
1- Put 2 in the sorted list and {3,1} in the unsorted list.
2- Compare 3 with 2 and do not swap since 2<3, add 3 to the sorted list so the sorted list is {2,3} and the unsorted list is {1}
3- Compare 1 with 3, and swap since 1<3 so the sorted list is {2,1,3} and unsorted list is empty. Continue by comparing 1 and 2 and swapping since 1<2 so the list becomes {1,2,3}
4- Since there are no elements left in the unsorted list, the list is sorted so return {1,2,3}
*/
func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
	}
	for i, test := range tests {
		InsertionSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
