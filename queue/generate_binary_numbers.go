package queue

import "errors"

type stringQueue struct {
	data []string
}

// GenerateBinaryNumbers returns string representation of binary numbers from 0 to n
func GenerateBinaryNumbers(n int) ([]string, error) {
	output := []string{"0"}
	if n == 0 {
		return output, nil
	}
	queue := new(stringQueue)

	queue.enqueue("1")
	for n > 0 {
		n--
		tmp, err := queue.dequeue()
		if err != nil {
			return nil, err
		}
		output = append(output, tmp)
		queue.enqueue(tmp + "0")
		queue.enqueue(tmp + "1")
	}

	return output, nil
}

func (q *stringQueue) len() int           { return len(q.data) }
func (q *stringQueue) enqueue(val string) { q.data = append(q.data, val) }
func (q *stringQueue) dequeue() (string, error) {
	if q.len() == 0 {
		return "", errors.New("Queue is empty")
	}
	tmp := q.data[0]
	q.data = q.data[1:len(q.data)]
	return tmp, nil
}
