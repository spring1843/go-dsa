package graph

// NumberOfIslands returns the number of islands or connected lands in a grid where 0 represents water
// and one represents land.
func NumberOfIslands(grid [][]int) int {
	islands := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				islands++
				visitBFS(grid, i, j)
			}
		}
	}
	return islands
}

func visitBFS(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != 1 {
		return
	}
	grid[i][j] = 2
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, direction := range directions {
		visitBFS(grid, i+direction[0], j+direction[1])
	}
}
