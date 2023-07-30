func knightProbability(n int, k int, row int, column int) float64 {
    dp := make([][][]float64, n)
    for i := range dp {
        dp[i] = make([][]float64, n)
        for j := range dp[i] {
            dp[i][j] = make([]float64, k+1)
            dp[i][j][0] = 1
        }
    }

    for round := 1; round <= k; round++ {
        for i := range dp {
            for j := range dp[i] {
                if i-1 >= 0 && j-2 >= 0 {
                    dp[i][j][round] += dp[i-1][j-2][round-1] * 0.125
                }

                if i-2 >= 0 && j-1 >= 0 {
                    dp[i][j][round] += dp[i-2][j-1][round-1] * 0.125
                }

                if i-2 >= 0 && j+1 < n {
                    dp[i][j][round] += dp[i-2][j+1][round-1] * 0.125
                }

                if i-1 >= 0 && j+2 < n {
                    dp[i][j][round] += dp[i-1][j+2][round-1] * 0.125
                }

                if i+1 < n && j+2 < n {
                    dp[i][j][round] += dp[i+1][j+2][round-1] * 0.125
                }

                if i+2 < n && j+1 < n {
                    dp[i][j][round] += dp[i+2][j+1][round-1] * 0.125
                }

                if i+2 < n && j-1 >= 0 {
                    dp[i][j][round] += dp[i+2][j-1][round-1] * 0.125
                }

                if i+1 < n && j-2 >= 0 {
                    dp[i][j][round] += dp[i+1][j-2][round-1] * 0.125
                }
            }
        }
    }

    return dp[row][column][k]
}

