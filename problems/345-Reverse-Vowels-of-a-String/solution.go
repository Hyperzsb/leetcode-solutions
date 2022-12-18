func isVowel(c byte) bool {
	if c == byte('a') || c == byte('A') || c == byte('e') || c == byte('E') ||
		c == byte('i') || c == byte('I') || c == byte('o') || c == byte('O') ||
		c == byte('u') || c == byte('U') {
		return true
	} else {
		return false
	}
}

func reverseVowels(s string) string {
	result := make([]byte, len(s))
	copy(result, s)

	start, end := 0, len(result)-1
	for start < end {
		if isVowel(result[start]) && isVowel(result[end]) {
			result[start] ^= result[end]
			result[end] ^= result[start]
			result[start] ^= result[end]

			start++
			end--
		}

		for start < len(result) && !isVowel(result[start]) {
			start++
		}

		for end >= 0 && !isVowel(result[end]) {
			end--
		}
	}

	return string(result)
}

