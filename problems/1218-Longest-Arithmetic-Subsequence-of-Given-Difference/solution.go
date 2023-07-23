func longestSubsequence(arr []int, difference int) int {
    dp:= make(map[int]int)
    for i := range arr {
        dp[arr[i]] = dp[arr[i]-difference] + 1
    }

    result := 0
    for _, length := range dp {
        if result < length {
            result = length
        }
    }

    return result
}

