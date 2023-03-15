package hashtable

// SumUpToK returns indices of the first pairs that sum up to k
func SumUpToK(numbers []int, k int) []int {
	output := []int{}
	hash := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if _, ok := hash[n]; ok {
			if n+n == k {
				return []int{hash[n], i}
			}
		}
		hash[n] = i
	}

	for i, n := range numbers {
		complement := k - n
		if _, ok := hash[complement]; ok {
			if hash[complement] != i {
				output = append(output, hash[complement])
			}
		}
	}
	if len(output) == 2 {
		output[0], output[1] = output[1], output[0]
	}
	return output
}
