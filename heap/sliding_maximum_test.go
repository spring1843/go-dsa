package heap

import (
	"container/heap"
	"reflect"
	"testing"
)

/*
TestMaxSlidingWindow tests solution(s) with the following signature and problem description:

	func MaxSlidingWindow(numbers []int, k int) []int {

Returns a maximum in each slice created by a sliding window of size k on a list of numbers.
*/
func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		numbers    []int
		k          int
		maxSliding []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1, 4, 5, -2, 4, 6}, 1, []int{1, 4, 5, -2, 4, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 2, []int{4, 5, 5, 4, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 3, []int{5, 5, 5, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 4, []int{5, 5, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 6, []int{6}},
	}

	for i, test := range tests {
		if got := MaxSlidingWindow(test.numbers, test.k); !reflect.DeepEqual(got, test.maxSliding) {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.maxSliding, got)
		}
	}
}

func TestMaxSlidingWindowPop(t *testing.T) {
	pq := make(slidingWindow, 5)
	heap.Init(&pq)
	heap.Push(&pq, 5)
	if got := heap.Pop(&pq).(int); got != 5 {
		t.Fatalf("Wanted %d got %d", got, 5)
	}
}
