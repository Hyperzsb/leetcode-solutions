func check(s string, sm *map[string]byte, count int) bool {
	if ((*sm)[s] == 1 && count > 0) || (*sm)[s] == 2 {
		return true
	}

	for i := 0; i < len(s)-1; i++ {
		if (*sm)[s[:i+1]] != 0 && check(s[i+1:], sm, count+1) {
			(*sm)[s] = 2
			return true
		}
	}

	return false
}

func findAllConcatenatedWordsInADict(words []string) []string {
	strMap := make(map[string]byte)
	for _, word := range words {
		strMap[word] = 1
	}

	result := make([]string, 0)
	for _, word := range words {
		if check(word, &strMap, 0) {
			result = append(result, word)
		}
	}

	return result
}

