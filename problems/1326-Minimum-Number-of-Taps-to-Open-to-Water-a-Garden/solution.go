func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func minTaps(n int, ranges []int) int {
    maxRight := make([]int, n+1)
    for i := range ranges {
        idx := max(0, i-ranges[i])
        maxRight[idx] = max(maxRight[idx], i+ranges[i])
    }

    end, right, result := 0, 0, 0
    for i := 0; i <= n; i++ {
        if i > end {
            if right <= end {
                return -1
            } else {
                end = right
                result++
            }
        }

        right = max(right, maxRight[i])
    }

    return result
}

