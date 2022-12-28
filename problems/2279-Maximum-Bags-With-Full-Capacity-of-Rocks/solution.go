func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}

	sort.Slice(capacity, func(a, b int) bool {
		return capacity[a] < capacity[b]
	})

	result := 0
	for i := 0; i < len(capacity); i++ {
		additionalRocks -= capacity[i]
		if additionalRocks >= 0 {
			result++
		} else {
			break
		}
	}

	return result
}

