func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(a, b int) bool {
		return costs[a] < costs[b]
	})

	result := 0
	for coins > 0 && result < len(costs) {
		coins -= costs[result]
		if coins >= 0 {
			result++
		}
	}

	return result
}

