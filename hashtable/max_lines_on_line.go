package hashtable

import "math"

// MaxPointsOnALine given a number of coordinates for points, returns the
// maximum number of points that are on the same line.
func MaxPointsOnALine(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	maxPoints := 1
	for i, point := range points {
		samePoint := 0
		localMax := 1
		slopeCount := make(map[float64]int)
		for j := i + 1; j < len(points); j++ {
			if isSame(point, points[j]) {
				samePoint++
				continue
			}
			slope := calculateSlope(point, points[j])
			slopeCount[slope]++
			localMax = max(localMax, slopeCount[slope])
		}
		maxPoints = max(maxPoints, localMax+samePoint+1)
	}
	return maxPoints
}

func isSame(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func calculateSlope(a, b []int) float64 {
	if a[0] == b[0] {
		return math.MaxFloat64
	}
	if a[1] == b[1] {
		return 1
	}
	return float64(a[1]-b[1]) / float64(a[0]-b[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
