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

func abs(n int) int {
    if n >= 0 {
        return n
    } else {
        return -n
    }
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
    nm := make(map[K]V)
    for k, v := range m {
        nm[k] = v
    }

    return nm
}

func tallestBillboard(rods []int) int {
    dp := make(map[int]int)
    dp[0] = 0

    for _, rod := range rods {
        ndp := copyMap(dp)
        for k, v := range ndp {
            dp[k+rod] = max(dp[k+rod], v)
            dp[abs(k-rod)] = max(dp[abs(k-rod)], v+min(k, rod))
        }
    }

    return dp[0]
}

