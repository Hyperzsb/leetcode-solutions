func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	dp := make([][]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]bool, sum/2+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= sum/2; j++ {
			if dp[i-1][j] || (nums[i-1] <= j && dp[i-1][j-nums[i-1]]) {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[len(nums)][sum/2]
}

