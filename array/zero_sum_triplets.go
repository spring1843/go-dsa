package array

import "sort"

func ZeroSumTriplets(list []int) [][]int {
	output := make([][]int, 0)
	if len(list) < 3 {
		return output
	}

	sort.Ints(list)
	for i, n := range list {
		if i > 0 && n == list[i-1] {
			continue
		}

		l, r := i+1, len(list)-1
		for l < r {
			threeSum := n + list[l] + list[r]
			if threeSum > 0 {
				r--
				continue
			}
			if threeSum < 0 {
				l++
				continue
			}
			output = append(output, []int{n, list[l], list[r]})
			l++
			for list[l] == list[l-1] && l < r {
				l++
			}
		}
	}
	return output
}
