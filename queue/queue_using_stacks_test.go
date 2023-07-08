package queue

import "testing"

/*
TestQueueUsingStacks tests solution(s) with the following signature and problem description:

	type UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	(usingStacks *UsingStacks) enqueue(n int)
	(usingStacks *UsingStacks) dequeue() int

Implement a queue of integers using stacks.
*/
func TestQueueUsingStacks(t *testing.T) {
	tests := []struct {
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
