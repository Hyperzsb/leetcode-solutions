func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	patternMap := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		if _, exist := patternMap[byte(pattern[i])]; exist {
			if patternMap[byte(pattern[i])] != strs[i] {
				return false
			}
		} else {
			for _, val := range patternMap {
				if val == strs[i] {
					return false
				}
			}
			patternMap[byte(pattern[i])] = strs[i]
		}
	}

	return true
}

