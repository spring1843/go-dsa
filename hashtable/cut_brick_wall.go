package hashtable

// CutBrickWall runs in O(n) time and O(n) space.
func CutBrickWall(wall [][]int) int {
	cuts := make(map[int]int)
	var max, position int
	for _, row := range wall {
		length := 0
		for i, brick := range row {
			if i == len(row)-1 {
				break
			}
			length += brick
			cuts[length]++
			if cuts[length] > max {
				max = cuts[length]
				position = length
			}
		}
	}
	return position
}
