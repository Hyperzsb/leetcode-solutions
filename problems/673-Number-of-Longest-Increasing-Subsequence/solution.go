func findNumberOfLIS(nums []int) int {
    dp, cnts := make([]int, len(nums)), make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
        cnts[i] = 1
    }

    longest := 1
    for i := range nums {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                    cnts[i] = cnts[j]
                    continue
                }

                if dp[j]+1 == dp[i] {
                    cnts[i] += cnts[j]
                }
            }
        }

        if longest < dp[i] {
            longest = dp[i]
        }
    }

    result := 0
    for i := range dp {
        if dp[i] == longest {
            result += cnts[i]
        }
    }

    return result
}

