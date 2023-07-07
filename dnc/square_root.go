package dnc

// SquareRoot solves the problem in O(log n) time and O(1) space.
func SquareRoot(number, precision int) float64 {
	start := 0
	end := number
	ans := 1.0

	for start <= end {
		mid := (start + end) / 2
		midSquare := mid * mid
		if midSquare == number {
			ans = float64(mid)
			break
		}

		if midSquare < number {
			start = mid + 1
			ans = float64(mid)
		} else {
			end = mid - 1
		}
	}

	increment := 0.1
	for i := 0; i < precision; i++ {
		for ans*ans <= float64(number) {
			ans += increment
		}
		ans -= increment
		increment /= 10
	}
	return ans
}
