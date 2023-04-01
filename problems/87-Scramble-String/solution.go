func isEqual(cm1, cm2 []int) bool {
	for i := 0; i < 26; i++ {
		if cm1[i] != cm2[i] {
			return false
		}
	}

	return true
}

func isS(s1, s2 string, history map[string]bool) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) == 1 {
		return false
	}

	key := s1 + ";" + s2
	if val, ok := history[key]; ok {
		return val
	}

	charMap1, charMap2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		charMap1[byte(s1[i])-'a']++
		charMap2[byte(s2[i])-'a']++
	}

	charMap11, charMap21 := make([]int, 26), make([]int, 26)
	charMap12, charMap22 := make([]int, 26), make([]int, 26)
	copy(charMap12, charMap1)
	copy(charMap22, charMap2)

	for i := 1; i < len(s1); i++ {
		charMap11[byte(s1[i-1])-'a']++
		charMap12[byte(s1[i-1])-'a']--
		charMap21[byte(s2[i-1])-'a']++
		charMap22[byte(s2[i-1])-'a']--

		if !isEqual(charMap11, charMap21) || !isEqual(charMap12, charMap22) {
			continue
		}

		if isS(string(s1[:i]), string(s2[:i]), history) && isS(string(s1[i:]), string(s2[i:]), history) {
			history[key] = true
			return true
		}
	}

	charMap11, charMap21 = make([]int, 26), make([]int, 26)
	charMap12, charMap22 = make([]int, 26), make([]int, 26)
	copy(charMap12, charMap1)
	copy(charMap22, charMap2)

	for i := 1; i < len(s1); i++ {
		charMap11[byte(s1[i-1])-'a']++
		charMap12[byte(s1[i-1])-'a']--
		charMap21[byte(s2[len(s2)-i])-'a']++
		charMap22[byte(s2[len(s2)-i])-'a']--

		if !isEqual(charMap11, charMap21) || !isEqual(charMap12, charMap22) {
			continue
		}

		if isS(string(s1[i:]), string(s2[:len(s2)-i]), history) && isS(string(s1[:i]), string(s2[len(s2)-i:]), history) {
			history[key] = true
			return true
		}
	}

	history[key] = false
	return false
}

func isScramble(s1 string, s2 string) bool {
	history := make(map[string]bool)
	return isS(s1, s2, history)
}

