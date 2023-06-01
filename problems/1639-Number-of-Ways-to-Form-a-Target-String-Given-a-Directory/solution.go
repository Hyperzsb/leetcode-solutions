func numWays(words []string, target string) int {
	slen, tlen := len(words[0]), len(target)

	charCnt := make([][]int, slen)
	for i := range charCnt {
		charCnt[i] = make([]int, 26)
		for j := range words {
			charCnt[i][words[j][i]-'a']++
		}
	}

	dp := make([][]int, tlen+1)
	for i := range dp {
		dp[i] = make([]int, slen+1)
	}
	dp[0][0] = 1

	for i := 1; i <= tlen; i++ {
		prevSum := 0
		for j := i; j <= slen; j++ {
			prevSum = (prevSum + dp[i-1][j-1]) % (1e9 + 7)

			cnt := charCnt[j-1][target[i-1]-'a']

			dp[i][j] = (prevSum * cnt) % (1e9 + 7)
		}
	}

	result := 0
	for i := tlen; i <= slen; i++ {
		result = (result + dp[tlen][i]) % (1e9 + 7)
	}

	return result
}

