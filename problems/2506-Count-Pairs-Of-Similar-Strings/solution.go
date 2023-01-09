func normalize(s string) string {
	charMap := make([]bool, 26)

	for _, char := range s {
		charMap[char-'a'] = true
	}

	result := make([]byte, 0, len(s))
	for char, exist := range charMap {
		if exist {
			result = append(result, byte(char+'a'))
		}
	}

	return string(result)
}

func similarPairs(words []string) int {
	patternMap := make(map[string]int)
	for _, word := range words {
		normalizedWord := normalize(word)
		patternMap[normalizedWord]++
	}

	result := 0
	for _, count := range patternMap {
		result += count * (count - 1) / 2
	}

	return result
}

