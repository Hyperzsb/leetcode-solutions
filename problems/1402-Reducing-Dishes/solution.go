func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(a, b int) bool {
		return satisfaction[a] > satisfaction[b]
	})

	result, prefixSum := 0, 0
	for i := 0; i < len(satisfaction); i++ {
		prefixSum += satisfaction[i]
		if prefixSum < 0 {
			break
		}

		result += prefixSum
	}

	return result
}

