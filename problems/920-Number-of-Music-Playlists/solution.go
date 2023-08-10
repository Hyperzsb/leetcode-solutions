func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func numMusicPlaylists(n int, goal int, k int) int {
    dp := make([][]int, goal+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    dp[0][0] = 1

    for i := 1; i <= goal; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = (dp[i-1][j-1] * (n - j + 1)) % (1e9 + 7)
            
            if j > k {
                dp[i][j] = (dp[i][j] + dp[i-1][j] * (j - k)) % (1e9 + 7)
            }
        }
    }

    fmt.Println(dp)

    return dp[goal][n]
}

