func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFlipsMonoIncr(s string) int {
	oneCount, flipCount := 0, 0

	for _, char := range s {
		if char == '1' {
			oneCount++
		} else {
			flipCount++
			flipCount = min(oneCount, flipCount)
		}
	}

	return flipCount
}

