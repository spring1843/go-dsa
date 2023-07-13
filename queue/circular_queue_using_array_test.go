package queue

import "testing"

type testCase struct {
	enqueue                  []int
	dequeueTimes             int
	expectedLastDequeuedItem int
	expectEnqueueErr         bool
	expectDequeueErr         bool
}

const (
	operationTypeEnqueue = iota
	operationTypeDequeue
)

var queueOperations = []struct {
	operationType  int
	operationValue int
}{
	{operationTypeEnqueue, 3},
	{operationTypeEnqueue, 2},
	{operationTypeEnqueue, 1},
	{operationTypeDequeue, 0},
	{operationTypeDequeue, 0},
	{operationTypeDequeue, 0},
	{operationTypeEnqueue, 3},
	{operationTypeEnqueue, 2},
	{operationTypeEnqueue, 1},
	{operationTypeDequeue, 0},
	{operationTypeDequeue, 0},
	{operationTypeDequeue, 0},
	{operationTypeEnqueue, 5},
	{operationTypeDequeue, 0},
	{operationTypeEnqueue, 1},
	{operationTypeDequeue, 0},
	{operationTypeEnqueue, 1},
}

/*
TestCircularQueue tests solution(s) with the following signature and problem description:

	func (queue *UsingCircularArray) enqueue(n int) error
	func (queue *UsingCircularArray) dequeue() (int, error)

Implements a queue using a circular array.
*/
func TestCircularQueue(t *testing.T) {
	// Tests a queue by enqueues given items,
	// then dequeues a number of times and finally checks the value from the last dequeue.
	// The same process then may be repeated for a second time with different values.
	tests := []struct {
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

func enqueueDequeueAndCheckValue(t *testing.T, queue *CircularQueue, testID int, test *testCase) {
	t.Helper()
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
	for i, queueOperation := range queueOperations {
		switch queueOperation.operationType {
		case operationTypeEnqueue:
			if err := queue.enqueue(queueOperation.operationValue); err != nil {
				t.Fatalf("operation #%d, unexpected error: %s", i, err)
			}
		case operationTypeDequeue:
			if _, err := queue.dequeue(); err != nil {
				t.Fatalf("operation #%d, unexpected error: %s", i, err)
			}
		}
	}

	n, err := queue.dequeue()
	if err != nil {
		t.Fatalf("Failed to dequeue. Error %s", err)
	}

	if n != 1 {
		t.Fatalf("Wanted %d, got %d", 1, n)
	}
}
