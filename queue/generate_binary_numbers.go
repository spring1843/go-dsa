package queue

type stringQueue struct {
	data []string
}

// GenerateBinaryNumbers solves the problem in O(n) time and O(n) space.
func GenerateBinaryNumbers(n int) []string {
	output := []string{"0"}
	if n == 0 {
		return output
	}
	queue := new(stringQueue)

	queue.enqueue("1")
	for n > 0 {
		n--
		tmp := queue.dequeue()
		output = append(output, tmp)
		queue.enqueue(tmp + "0")
		queue.enqueue(tmp + "1")
	}

	return output
}

func (q *stringQueue) enqueue(val string) { q.data = append(q.data, val) }
func (q *stringQueue) dequeue() string {
	tmp := q.data[0]
	q.data = q.data[1:len(q.data)]
	return tmp
}
