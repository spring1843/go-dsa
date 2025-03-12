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
	var i int
	for i = 0; i < k; i++ {
		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
	}

	for i < len(numbers) {
		output = append(output, numbers[queue.Front().Value.(int)])
		for queue.Len() != 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}

		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
		i++
	}

	output = append(output, numbers[queue.Front().Value.(int)])
	return output, nil
}
