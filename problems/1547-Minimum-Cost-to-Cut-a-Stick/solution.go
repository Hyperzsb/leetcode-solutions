func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func cut(left, right int, cuts []int, resultMap map[string]int) int {
	key := fmt.Sprintf("%d:%d", left, right)

	if val, ok := resultMap[key]; ok {
		return val
	}

	if len(cuts) == 0 {
		resultMap[key] = 0
		return 0
	}

	minCost := int(1e9)
	for i, c := range cuts {
		minCost = min(minCost,
			cut(left, c, cuts[:i], resultMap)+
				cut(c, right, cuts[i+1:], resultMap))
	}

	resultMap[key] = minCost + right - left
	return minCost + right - left
}

func minCost(n int, cuts []int) int {
	sort.Slice(cuts, func(a, b int) bool {
		return cuts[a] < cuts[b]
	})
	resultMap := make(map[string]int)

	return cut(0, n, cuts, resultMap)
}

