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

Implement a queue of integers using two stacks. The queue should support enqueue and dequeue operations.
*/
func TestQueueUsingStacks(t *testing.T) {
	tests := []struct {
		enqueue                  []int
		dequeueTimes             int
		expectedLastDequeuedItem int
	}{
		{[]int{1, 2, 3, 4}, 1, 1},
		{[]int{1, 2, 3, 4}, 2, 2},
		{[]int{1, 2, 3, 4}, 3, 3},
		{[]int{1, 2, 2, 3}, 2, 2},
	}

	for i, test := range tests {
		queue := new(UsingStacks)
		for _, n := range test.enqueue {
			queue.enqueue(n)
		}

		var n int
		for range test.dequeueTimes {
			n = queue.dequeue()
		}
		if n != test.expectedLastDequeuedItem {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.expectedLastDequeuedItem, n)
		}
	}
}
