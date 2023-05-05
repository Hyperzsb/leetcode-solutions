func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	} else {
		return false
	}
}

func maxVowels(s string, k int) int {
	vowelCnt := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			vowelCnt++
		}
	}

	result := vowelCnt

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			vowelCnt--
		}

		if isVowel(s[i]) {
			vowelCnt++

			if result < vowelCnt {
				result = vowelCnt
			}
		}
	}

	return result
}

