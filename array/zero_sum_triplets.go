package array

import "sort"

// ZeroSumTriplets solves the problem in O(n^2) time and O(n) space.
func ZeroSumTriplets(list []int) [][]int {
	res := make([][]int, 0)

	if len(list) < 3 {
		return res
	}

	sort.Ints(list)

	for i := range list {
		l := i + 1
		r := len(list) - 1

		for l < r {
			sum := list[i] + list[l] + list[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{list[i], list[l], list[r]})
				break
			}
		}
	}

	return res
}
