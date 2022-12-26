func findWords(words []string) []string {
	rows := make([]string, 3)
	rows[0] = "qwertyuiop"
	rows[1] = "asdfghjkl"
	rows[2] = "zxcvbnm"

	keys := make(map[byte]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < len(rows[i]); j++ {
			keys[rows[i][j]] = i
			keys[rows[i][j]-32] = i
		}
	}

	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		currentRow := keys[words[i][0]]
		flag := true
		for j := 1; j < len(words[i]); j++ {
			if keys[words[i][j]] != currentRow {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, words[i])
		}
	}

	return result
}

