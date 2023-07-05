package graph

import (
	"math"
	"testing"
)

func TestCheapestFlights(t *testing.T) {
	tests := []struct {
		flight                                                 [][]int
		cityCount, source, destination, maxStops, cheapestCost int
	}{
		{[][]int{{0, 1, math.MaxInt64}}, 3, 0, 2, 2, -1},
		{[][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 5}}, 3, 0, 2, 2, 5},
		{[][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 5}}, 3, 0, 2, 2, 5},
		{[][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 10}}, 3, 0, 2, 2, 10},
		{[][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 20}}, 3, 0, 2, 1, 20},
		{[][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 20}}, 3, 0, 2, 1, 20},
	}

	for i, test := range tests {
		got := CheapestFlights(test.flight, test.cityCount, test.source, test.destination, test.maxStops)
		if got != test.cheapestCost {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.cheapestCost, got)
		}
	}
}
