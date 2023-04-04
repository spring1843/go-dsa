package queue

import "testing"

type testCase struct {
	enqueue                  []int
	dequeueTimes             int
	expectedLastDequeuedItem int
	expectEnqueueErr         bool
	expectDequeueErr         bool
}

// TestCircularQueue tests a queue by enqueues given items,
// then dequeues a number of times and finally checks the value from the last dequeue.
// The same process then may be repeated for a second time with different values.
func TestCircularQueue(t *testing.T) {
	var tests = []struct {
		size        int
		firstRound  *testCase
		secondRound *testCase
	}{
		{0, &testCase{[]int{1, 2, 3, 4}, -1, -1, true, false}, nil},
		{4, &testCase{[]int{1, 2, 3, 4}, 1, 1, false, false}, nil},
		{4, &testCase{[]int{1, 2, 3, 4}, 2, 2, false, false}, nil},
		{4, &testCase{[]int{1, 2, 3, 4}, 3, 3, false, false}, nil},
		{4, &testCase{[]int{1, 2, 2, 3}, 2, 2, false, false}, nil},
		{2, &testCase{[]int{1, 2, 2, 3}, 2, 2, true, false}, nil},
		{2, &testCase{[]int{1, 2}, 3, -1, false, true}, nil},
		{2, &testCase{[]int{1, 2}, 2, 2, false, false}, &testCase{[]int{1, 2}, 1, 1, false, false}},
		{2, &testCase{[]int{1, 2}, 1, 1, false, false}, &testCase{[]int{1}, 2, 1, false, false}},
	}

	for i, test := range tests {
		queue := NewCircularQueue(test.size)

		enqueueDequeueAndCheckValue(t, queue, i, test.firstRound)
		if test.secondRound != nil {
			enqueueDequeueAndCheckValue(t, queue, i, test.secondRound)
		}
	}
}

func enqueueDequeueAndCheckValue(t *testing.T, queue *UsingCircularArray, testID int, test *testCase) {
	for _, n := range test.enqueue {
		if err := queue.enqueue(n); err != nil {
			if test.expectEnqueueErr {
				t.Skipf("Skipping test case #%d. Expected error occurred. Error %s", testID, err)
			}
			t.Fatalf("Failed test case #%d. Did not expect enqueuing error. Error %s", testID, err)
		}
	}

	var n int
	var err error
	for j := 1; j <= test.dequeueTimes; j++ {
		n, err = queue.dequeue()
		if err != nil {
			if test.expectDequeueErr {
				t.Skipf("Skipping test case #%d. Expected error occurred. Error %s", testID, err)
			}
			t.Fatalf("Failed test case #%d. Did not expect dequeuing error. Error %s", testID, err)
		}
	}
	if n != test.expectedLastDequeuedItem {
		t.Fatalf("Failed test case #%d. Want %d got %d", testID, test.expectedLastDequeuedItem, n)
	}
}

func TestMultipleOperations(t *testing.T) {
	queue := NewCircularQueue(3)
	if err := queue.enqueue(3); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(2); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(1); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(3); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(2); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(1); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(5); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(1); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if _, err := queue.dequeue(); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if err := queue.enqueue(1); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	n, err := queue.dequeue()
	if err != nil {
		t.Fatalf("Failed dequing. Error %s", err)
	}

	if n != 1 {
		t.Fatalf("Wanted %d, got %d", 1, n)
	}
}
