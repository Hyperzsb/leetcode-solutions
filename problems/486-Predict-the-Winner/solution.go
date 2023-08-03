func maxScore(nums []int, s, e int, sums [][]int, dp [][]int) int {
    if dp[s][e] != -1 {
        return dp[s][e]
    }

    pickHead := maxScore(nums, s+1, e, sums, dp)
    pickTail := maxScore(nums, s, e-1, sums, dp)
    
    if pickHead >= pickTail {
        dp[s][e] = sums[s][e] - pickTail
    } else {
        dp[s][e] = sums[s][e] - pickHead
    }

    return dp[s][e]
}

func PredictTheWinner(nums []int) bool {
    sums := make([][]int, len(nums))
    for i := range sums {
        sums[i] = make([]int, len(nums))

        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            sums[i][j] = sum
        }
    }

    dp := make([][]int, len(nums))
    for i := range dp {
        dp[i] = make([]int, len(nums))
        
        for j := range dp[i] {
            dp[i][j] = -1
        }

        dp[i][i] = nums[i]
    }

    score := maxScore(nums, 0, len(nums)-1, sums, dp)
    if score >= sums[0][len(nums)-1]-score {
        return true
    } else {
        return false
    }
}

