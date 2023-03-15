package array

import "math"

// AddTwoNumbers adds two numbers which are represented as an array and returns the results
// In the same format. For example [2,5], and [3,5] would add up to[6,0] because 25 + 35 = 60.
func AddTwoNumbers(num1, num2 []int) []int {
	num1, num2 = equalizeLengths(num1, num2)
	carry := false
	for i := len(num1) - 1; i > -1; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}
	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

func equalizeLengths(num1, num2 []int) ([]int, []int) {
	diff := int(math.Abs(float64(len(num2) - len(num1))))
	zeros := make([]int, diff)
	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)
	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}
