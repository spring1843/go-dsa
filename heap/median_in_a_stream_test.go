package heap

import "testing"

func TestMedianInAStream(t *testing.T) {
	tests := []struct {
		numbers []int
		median  float64
	}{
		{[]int{1}, 1},
		{[]int{2, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{6, 5, 4, 3, 2, 1}, 3.5},
		{[]int{1, 4, 5, -2, 4, 6}, 4},
	}

	for i, test := range tests {
		heap := newMedianKeeper()
		for _, number := range test.numbers {
			heap.addNumber(number)
		}
		got := heap.median()
		if got != test.median {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.median, got)
		}
	}
}
