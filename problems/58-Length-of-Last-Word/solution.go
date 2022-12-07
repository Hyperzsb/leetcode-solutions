func lengthOfLastWord(s string) int {
	length, flag := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && flag == 1 {
			flag = 0
			continue
		} else if s[i] != ' ' && flag == 0 {
			flag = 1
			length = 1
		} else if s[i] != ' ' && flag == 1 {
			length++
		} else {
			continue
		}
	}

	return length
}

