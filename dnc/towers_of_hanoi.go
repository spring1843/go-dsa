package dnc

type moves [][2]int

// TowerOfHanoi will return the moves it takes to move all disks from start to end with respect
// to the rules of the Tower of Hanoi game where:
// n is the number of disks stacked on top of each other
// heavier disks can never be placed on a lighter disk
// There are 3 towers, all disks are initially places at the start tower.
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
