package graph

import (
	"container/list"
	"strconv"
	"strings"
)

const (
	lineSeparator   = "\n"
	numberSeparator = ","
)

var errInvalidInteger error

// HeadCount solves the problem in O(n) time and O(n) space.
func HeadCount(data string, employeeID int) int {
	graph, err := toGraphOfEmployees(data)
	if err != nil {
		return -1
	}

	bfs := func(employeeID int) int {
		reportCount := 0
		queue := list.New()
		queue.PushBack(employeeID)

		for queue.Len() != 0 {
			employee := queue.Back().Value.(int)
			queue.Remove(queue.Back())

			if graph == nil {
				return -1
			}

			for _, report := range graph[employee] {
				queue.PushBack(report)
				reportCount++
			}
		}
		return reportCount + 1
	}
	return bfs(employeeID)
}

// toGraphOfEmployees creates a map of employees and their direct reports.
func toGraphOfEmployees(data string) (map[int][]int, error) {
	output := make(map[int][]int)
	for _, line := range strings.Split(data, lineSeparator) {
		broken := make([]int, 0)
		for _, item := range strings.Split(line, numberSeparator) {
			n, err := strconv.Atoi(item)
			if err != nil {
				return nil, errInvalidInteger
			}
			broken = append(broken, n)
		}
		output[broken[0]] = broken[1:]
	}
	return output, nil
}
