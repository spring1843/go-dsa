package backtracking

var (
	directions      = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directionLetter = map[int]string{0: "r", 1: "l", 2: "d", 3: "u"}
)

// Maze solves the problem in O(mn) time and O(mn) space complexity.
func Maze(m, n int, walls [][2]int, start, finish [2]int) string {
	wallMap := make(map[[2]int]bool)
	for _, wall := range walls {
		wallMap[wall] = true
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return mazeRecursive(start[0], start[1], "", wallMap, visited, finish)
}

func mazeRecursive(x, y int, path string, wallMap map[[2]int]bool, visited [][]bool, finish [2]int) string {
	if x == finish[0] && y == finish[1] {
		return path
	}

	visited[x][y] = true
	for i := range 4 {
		nx, ny := x+directions[i][0], y+directions[i][1]
		if canVisit(nx, ny, wallMap, visited) {
			if result := mazeRecursive(nx, ny, path+directionLetter[i], wallMap, visited, finish); result != "" {
				return result
			}
		}
	}
	return ""
}

func canVisit(x, y int, wallMap map[[2]int]bool, visited [][]bool) bool {
	inBounds := x >= 0 && x < len(visited) && y >= 0 && y < len(visited[0])
	return inBounds && !visited[x][y] && !wallMap[[2]int{x, y}]
}
