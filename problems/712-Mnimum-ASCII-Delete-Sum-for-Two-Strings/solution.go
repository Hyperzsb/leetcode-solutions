func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func traverse(s1, s2 string, idx1, idx2 int, dp [][]int) int {
    if idx1 == len(s1) || idx2 == len(s2) {
        return 0
    }

    if dp[idx1][idx2] != -1 {
        return dp[idx1][idx2]
    }

    if s1[idx1] == s2[idx2] {
        dp[idx1][idx2] = int(s1[idx1]) + traverse(s1, s2, idx1+1, idx2+1, dp)
    } else {
        dp[idx1][idx2] = max(traverse(s1, s2, idx1+1, idx2, dp), traverse(s1, s2, idx1, idx2+1, dp))
    }

    return dp[idx1][idx2]
}

func minimumDeleteSum(s1 string, s2 string) int {
    sum := 0

    for i := range s1 {
        sum += int(s1[i])
    }

    for i := range s2 {
        sum += int(s2[i])
    }

    dp := make([][]int, len(s1))
    for i := range dp {
        dp[i] = make([]int, len(s2))
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    return sum - 2*traverse(s1, s2, 0, 0, dp)
}

