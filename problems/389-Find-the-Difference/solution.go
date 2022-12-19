func findTheDifference(s string, t string) byte {
	charCount := make([]int, 26)
	for _, val := range s {
		charCount[val-'a']++
	}

	var result byte
	for _, val := range t {
		charCount[val-'a']--
		if charCount[val-'a'] == -1 {
			result = byte(val)
		}
	}

	return result
}

