package graph

import "math"

// CheapestFlights finds cheapest flights with up to k stops from vertexSource to vertexDestination given a collection of flights
// and their costs
func CheapestFlights(flights [][]int, cityCount, source, destination, maxStops int) int {
	var (
		min        = math.MaxInt64
		graph      = flightGraph(flights, cityCount)
		findMinDFS func(vertex, distance, depth int)
	)

	findMinDFS = func(vertex, distance, depth int) {
		if depth > maxStops {
			return
		}
		if vertex == destination && distance < min {
			min = distance
		}
		for _, edge := range graph[vertex] {
			if temp := distance + edge[1]; temp < min {
				findMinDFS(edge[0], temp, depth+1)
			}
		}
	}

	findMinDFS(source, 0, 0)
	if min == math.MaxInt64 {
		return -1
	}
	return min
}

func flightGraph(flights [][]int, cityCount int) [][][]int {
	graph := make([][][]int, cityCount)
	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], flight[1:])
	}
	return graph
}
