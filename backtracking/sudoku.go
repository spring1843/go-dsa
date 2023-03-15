package backtracking

// Sudoku solves a given partially filled or empty 9x9 Soduku board by placing
// integers between 1 and 9 in empty spot designated by 0 such that
// In each row, column, and 3x3 sub square the values are unique.
func Sudoku(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 0 {
				continue
			}
			for value := 1; value <= 9; value++ {
				if !isValidSudokuPlacement(board, i, j, value) {
					continue
				}
				board[i][j] = value
				if Sudoku(board) {
					return true
				}
				board[i][j] = 0
			}
			return false
		}
	}
	return true
}

func isValidSudokuPlacement(board [][]int, row, col, value int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == value || board[row][i] == value || board[3*(row/3)+i/3][3*(col/3)+i%3] == value {
			return false
		}
	}
	return true
}
