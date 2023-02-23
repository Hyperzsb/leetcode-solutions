func isCapable(weights *[]int, days, capacity int) bool {
	idx := 0
	for idx < len(*weights) {
		if days == 0 {
			return false
		}

		weight := 0
		for idx < len(*weights) && weight+(*weights)[idx] <= capacity {
			weight += (*weights)[idx]
			idx++
		}

		if weight == 0 {
			return false
		}

		days--
	}

	return true
}

func roundupDiv(a, b int) int {
	if a%b == 0 {
		return a / b
	} else {
		return a/b + 1
	}
}

func shipWithinDays(weights []int, days int) int {
	minWeight, maxWeight := 500, 1
	for i := 0; i < len(weights); i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}

		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	minCap, maxCap := minWeight*roundupDiv(len(weights), days), maxWeight*roundupDiv(len(weights), days)
	halfCap := (minCap + maxCap) / 2

	for minCap < maxCap {
		if isCapable(&weights, days, halfCap) {
			maxCap = halfCap
		} else {
			minCap = halfCap + 1
		}

		halfCap = (minCap + maxCap) / 2
	}

	return halfCap
}

