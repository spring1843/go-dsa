package dnc

// SquareRoot solves the problem in O(Log n) time and O(1) space.
func SquareRoot(number, precision int) float64 {
	start := 0
	end := number
	answer := 1.0

	for start <= end {
		mid := (start + end) / 2
		midSquare := mid * mid
		if midSquare == number {
			answer = float64(mid)
			break
		}

		if midSquare < number {
			start = mid + 1
			answer = float64(mid)
		} else {
			end = mid - 1
		}
	}

	increment := 0.1
	for range precision {
		for answer*answer <= float64(number) {
			answer += increment
		}
		answer -= increment
		increment /= 10
	}
	return answer
}
