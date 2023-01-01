func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	chars1, chars2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(word1); i++ {
		chars1[word1[i]-'a']++
		chars2[word2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if (chars1[i] == 0 && chars2[i] != 0) ||
			(chars1[i] != 0 && chars2[i] == 0) {
			return false
		}
	}

	sort.Slice(chars1, func(a, b int) bool {
		return chars1[a] < chars1[b]
	})
	sort.Slice(chars2, func(a, b int) bool {
		return chars2[a] < chars2[b]
	})

	for i := 0; i < 26; i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	return true
}

