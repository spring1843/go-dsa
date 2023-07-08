package graph

import "testing"

/*
TestNetworkDelayTime tests solution(s) with the following signature and problem description:

	func NetworkDelayTime(n, k int, edges [][3]int) int

Given `n`, the number of nodes in a network, a list of unidirectional travel times in
`{source, destination, delay}` format, and `k`, an origin node number, return the time it will
take for a message sent from k to be received by all other nodes.
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
