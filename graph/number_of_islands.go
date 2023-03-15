package graph

// NumberOfIslands returns the number of islands or connected lands in a grid where 0 represents water
// and one represents land
func NumberOfIslands(grid [][]int) int {
	islands := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				islands++
				grid = visitBFS(grid, i, j)
			}
		}
	}
	return islands
}

func visitBFS(grid [][]int, i, j int) [][]int {
	if grid[i][j] != 1 {
		return grid
	}
	grid[i][j] = 2

	if i+1 < len(grid) {
		grid = visitBFS(grid, i+1, j)
	}
	if i-1 >= 0 {
		grid = visitBFS(grid, i-1, j)
	}
	if len(grid) > 0 && j+1 < len(grid[0]) {
		grid = visitBFS(grid, i, j+1)
	}
	if j-1 >= 0 {
		grid = visitBFS(grid, i, j-1)
	}
	return grid
}
