func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0)

	idx1, idx2 := 0, 0
	turn := 1
	for idx1 < len(word1) && idx2 < len(word2) {
		if turn == 1 {
			result = append(result, word1[idx1])
			idx1++
			turn = 2
		} else {
			result = append(result, word2[idx2])
			idx2++
			turn = 1
		}
	}

	result = append(result, word1[idx1:]...)
	result = append(result, word2[idx2:]...)

	return string(result)
}

