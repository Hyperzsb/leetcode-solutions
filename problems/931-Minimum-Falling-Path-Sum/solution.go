func minFallingPathSum(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])

	dp := make([][]int, row)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < col; i++ {
		dp[row-1][i] = matrix[row-1][i]
	}

	for i := row - 2; i >= 0; i-- {
		for j := 0; j < col; j++ {
			if j == 0 {
				if dp[i+1][j] <= dp[i+1][j+1] {
					dp[i][j] = matrix[i][j] + dp[i+1][j]
				} else {
					dp[i][j] = matrix[i][j] + dp[i+1][j+1]
				}
			} else if j == col-1 {
				if dp[i+1][j-1] <= dp[i+1][j] {
					dp[i][j] = matrix[i][j] + dp[i+1][j-1]
				} else {
					dp[i][j] = matrix[i][j] + dp[i+1][j]
				}
			} else {
				if dp[i+1][j-1] <= dp[i+1][j] && dp[i+1][j-1] <= dp[i+1][j+1] {
					dp[i][j] = matrix[i][j] + dp[i+1][j-1]
				} else if dp[i+1][j] <= dp[i+1][j-1] && dp[i+1][j] <= dp[i+1][j+1] {
					dp[i][j] = matrix[i][j] + dp[i+1][j]
				} else {
					dp[i][j] = matrix[i][j] + dp[i+1][j+1]
				}
			}
		}
	}

	min := 1000000
	for i := 0; i < col; i++ {
		if dp[0][i] <= min {
			min = dp[0][i]
		}
	}

	return min
}

