func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func integerBreak(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1

    for i := 2; i <= n; i++ {
        for j := 1; j < i; j++ {
            dp[i] = max(dp[i], max(dp[i-j], i-j) * max(dp[j], j))
        }
    }

    return dp[n]
}

