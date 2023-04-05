package graph

import (
	"container/list"
	"strconv"
	"strings"
)

type HeadCounter struct {
	data map[int][]int
}

const (
	lineSeparator   = "\n"
	numberSeparator = ","
)

// HeadCount returns 1 + the number of directs, and all their directs for a given employee.
func (h *HeadCounter) HeadCount(employeeID int) int {
	bfs := func(employeeID int) int {
		reportCount := 0
		queue := list.New()
		queue.PushBack(employeeID)

		for queue.Len() != 0 {
			employee := queue.Back().Value.(int)
			queue.Remove(queue.Back())

			for _, report := range h.data[employee] {
				queue.PushBack(report)
				reportCount++
			}
		}
		return reportCount + 1
	}
	return bfs(employeeID)
}

// NewHeadCount returns a head counter with initialized data.
func NewHeadCount(data string) (*HeadCounter, error) {
	headCounter := &HeadCounter{
		data: make(map[int][]int),
	}
	for _, line := range strings.Split(data, lineSeparator) {
		broken := make([]int, 0)
		for _, item := range strings.Split(line, numberSeparator) {
			n, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
			broken = append(broken, n)
		}
		headCounter.data[broken[0]] = broken[1:]
	}
	return headCounter, nil
}
