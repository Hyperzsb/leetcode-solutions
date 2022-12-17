func longestPalindrome(s string) string {
	checkList := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		checkList[i] = make([]int, 2)
		checkList[i][0] = 0
		checkList[i][1] = 1
	}

	max := 1
	result := string(s[0])
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(checkList[i-1]); j++ {
			if i-1-checkList[i-1][j] >= 0 && s[i-1-checkList[i-1][j]] == s[i] {
				checkList[i] = append(checkList[i], checkList[i-1][j]+2)
				if checkList[i-1][j]+2 > max {
					max = checkList[i-1][j] + 2
					result = s[i-1-checkList[i-1][j] : i+1]
				}
			}
		}
	}

	return result
}

