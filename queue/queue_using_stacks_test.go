package queue

import "testing"

func TestQueueUsingStacks(t *testing.T) {
	var tests = []struct {
		enqueue                  []int
		dequeueTimes             int
		expectedLastDequeuedItem int
		expectErr                bool
	}{
		{[]int{1, 2, 3, 4}, 1, 1, false},
		{[]int{1, 2, 3, 4}, 2, 2, false},
		{[]int{1, 2, 3, 4}, 3, 3, false},
		{[]int{1, 2, 2, 3}, 2, 2, false},
	}

	for i, test := range tests {
		queue := new(UsingStacks)
		for _, n := range test.enqueue {
			queue.enqueue(n)
		}

		var n int
		var err error
		for j := 1; j <= test.dequeueTimes; j++ {
			n = queue.dequeue()
			if !test.expectErr && err != nil {
				t.Fatalf("Failed test case #%d. Did not expect an error. Error %s", i, err)
			}
		}
		if n != test.expectedLastDequeuedItem {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.expectedLastDequeuedItem, n)
		}
	}
}
