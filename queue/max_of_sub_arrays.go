package queue

import (
	"container/list"
	"errors"
)

// ErrQueueEmpty occurs when dequeuing an empty queue.
var ErrQueueEmpty = errors.New("empty queue")

// MaxOfKLengthSubArrays solves the problem in O(n) time and O(k) space.
func MaxOfKLengthSubArrays(numbers []int, k int) ([]int, error) {
	output := []int{}
	queue := list.New()
	for i := range k {
		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
	}

	for i := k; i < len(numbers); i++ {
		output = append(output, numbers[queue.Front().Value.(int)])
		for queue.Len() != 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}

		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
	}

	output = append(output, numbers[queue.Front().Value.(int)])
	return output, nil
}
