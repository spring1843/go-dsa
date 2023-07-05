package dp

import (
	"testing"
)

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
