package bit

// Divide divides two numbers without using division or multiplication symbols
func Divide(x, y int) int {
	if y == 0 {
		return 0
	}

	quotient := 0
	power := uint64(32)
	yPower := y << power
	reminder := x

	for reminder >= y {
		for yPower > reminder {
			yPower >>= 1
			power--
		}
		quotient += 1 << power
		reminder -= yPower
	}
	return quotient
}
