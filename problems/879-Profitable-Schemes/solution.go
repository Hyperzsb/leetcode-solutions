func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func backtrace(gro, pro, crime, n, minProfit int, group, profit []int, resultMap [][][]int) int {
	if resultMap[gro][pro][crime] != -1 {
		return resultMap[gro][pro][crime]
	}

	result := 0
	if crime == len(profit) {
		if pro >= minProfit {
			result = 1
		} else {
			result = 0
		}

		resultMap[gro][pro][crime] = result

		return result
	}

	if gro+group[crime] <= n {
		result = backtrace(gro+group[crime], min(minProfit, pro+profit[crime]), crime+1, n, minProfit, group, profit, resultMap)
	}
	result = (result + backtrace(gro, pro, crime+1, n, minProfit, group, profit, resultMap)) % (1e9 + 7)

	resultMap[gro][pro][crime] = result

	return result
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	resultMap := make([][][]int, n+1)
	for i := range resultMap {
		resultMap[i] = make([][]int, minProfit+1)
		for j := range resultMap[i] {
			resultMap[i][j] = make([]int, len(profit)+1)
			for k := range resultMap[i][j] {
				resultMap[i][j][k] = -1
			}
		}
	}

	return backtrace(0, 0, 0, n, minProfit, group, profit, resultMap)
}

