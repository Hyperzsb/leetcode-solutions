func reset(charMap []bool) {
	for i := 0; i < 26; i++ {
		charMap[i] = false
	}
}

func partitionString(s string) int {
	charMap := make([]bool, 26)

	result := 0
	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'] {
			result++
			reset(charMap)
		}

		charMap[s[i]-'a'] = true
	}

	return result + 1
}

