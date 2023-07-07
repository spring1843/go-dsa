package graph

import "testing"

/*
TestNetworkDelayTime tests solution(s) with the following signature and problem description:

	func NetworkDelayTime(n, k int, edges [][3]int) int {

Returns the fastest times it will take for a message sent from k
to reach all nodes of the given graph of n vertices. The edges are given in the
form of {vertexSource, vertexDestination, vertexDelay}.
*/
func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		edges [][3]int // source, destination, delay
		numberOfNodes,
		messageFrom,
		shortestPath int
	}{
		{[][3]int{{0, 1, 1}}, 2, 1, -1},
		{[][3]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 0, 5},
		{[][3]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 1, -1},
		{[][3]int{{0, 1, 2}, {0, 2, 3}, {2, 4, 2}, {1, 3, 2}, {3, 4, 3}, {4, 5, 2}}, 5, 0, 7},
		{[][3]int{{0, 1, 2}, {0, 2, 3}, {2, 4, 2}, {1, 3, 2}, {3, 4, 3}, {4, 5, 2}, {5, 0, 1}}, 5, 5, 5},
	}

	for i, test := range tests {
		if got := NetworkDelayTime(test.numberOfNodes, test.messageFrom, test.edges); got != test.shortestPath {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.shortestPath, got)
		}
	}
}
