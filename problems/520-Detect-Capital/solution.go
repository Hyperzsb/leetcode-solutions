func check(s string) bool {
	allUpper, allLower := true, true
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			allUpper = false
		} else {
			allLower = false
		}
	}

	return allUpper || allLower
}

func detectCapitalUse(word string) bool {
	if word[0] >= 'A' && word[0] <= 'Z' {
		return check(string(word[1:]))
	}

	return check(word)
}

