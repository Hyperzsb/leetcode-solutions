func findKthPositive(arr []int, k int) int {
	result := 1
	for _, val := range arr {
		if val == result {
			result++
			continue
		}

		if skip := val - result; k <= skip {
			return result + k - 1
		} else {
			result += skip
			k -= skip
		}

		result++
	}

	return result + k - 1
}

