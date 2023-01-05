func minDeletionSize(strs []string) int {
	result := 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if j > 0 && strs[j][i] < strs[j-1][i] {
				result++
				break
			}
		}
	}

	return result
}

