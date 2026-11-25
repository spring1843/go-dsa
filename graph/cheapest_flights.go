package graph

import "math"

// CheapestFlights solves the problem in O(n^2) time and O(n) space.
func CheapestFlights(flights [][]int, cityCount, source, destination, maxStops int) int {
	var (
		cheapest   = math.MaxInt64
		graph      = flightGraph(flights, cityCount)
		findMinDFS func(vertex, distance, depth int)
	)

	findMinDFS = func(vertex, distance, depth int) {
		if depth > maxStops+1 {
			return
		}
		if vertex == destination && distance < cheapest {
			cheapest = distance
		}
		for _, edge := range graph[vertex] {
			if temp := distance + edge[1]; temp < cheapest {
				findMinDFS(edge[0], temp, depth+1)
			}
		}
	}

	findMinDFS(source, 0, 0)
	if cheapest == math.MaxInt64 {
		return -1
	}
	return cheapest
}

func flightGraph(flights [][]int, cityCount int) [][][]int {
	graph := make([][][]int, cityCount)
	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], flight[1:])
	}
	return graph
}
