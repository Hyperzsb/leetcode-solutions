func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1)+len(s2) != len(s3) {
        return false
    }

    if s1 == s3 || s2 == s3 {
        return true
    }

    dp := make([][]bool, len(s1)+1)
    for i := range dp {
        dp[i] = make([]bool, len(s2)+1)
    }

    dp[0][0] = true

    for i := 1; i <= len(s3); i++ {
        for j := max(0, i-1-len(s2)); j <= min(i-1, len(s1)); j++ {
            if dp[j][i-1-j] {
                if j < len(s1) && s1[j] == s3[i-1] {
                    dp[j+1][i-1-j] = true
                } 

                if i-1-j < len(s2) && s2[i-1-j] == s3[i-1] {
                    dp[j][i-j] = true
                }
            }
        }
    }

    return dp[len(s1)][len(s2)]
}

