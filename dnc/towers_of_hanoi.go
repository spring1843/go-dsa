package dnc

type moves [][2]int

// TowerOfHanoi solves the problem in O(2^n) time and O(n) space.
func TowerOfHanoi(n, start, end int) [][2]int {
	m := make(moves, 0)
	towerOfHanoiRecursive(n, start, end, &m)
	return m
}

func towerOfHanoiRecursive(n, start, end int, moves *moves) {
	if n <= 1 {
		*moves = append(*moves, [2]int{start, end})
		return
	}
	other := 6 - (start + end)
	towerOfHanoiRecursive(n-1, start, other, moves)
	*moves = append(*moves, [2]int{start, end})
	towerOfHanoiRecursive(n-1, other, end, moves)
}
