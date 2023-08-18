func check(flags []bool) bool {
    for i := range flags {
        if flags[i] {
            return true
        }
    }

    return false
}

func validPartition(nums []int) bool {
    if len(nums) == 2 {
        return nums[0] == nums[1]
    }

    dp := make([][]bool, len(nums))
    for i := range dp {
        dp[i] = make([]bool, 3)
    }

    if nums[0] == nums[1] {
        dp[1][0] = true
    }

    if nums[0] == nums[1] && nums[1] == nums[2] {
        dp[2][1] = true
    }

    if nums[0]+1 == nums[1] && nums[1]+1 == nums[2] {
        dp[2][2] = true
    }

    for i := 3; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            if check(dp[i-2]) {
                dp[i][0] = true
            }
        }

        if nums[i] == nums[i-1] && nums[i-1] == nums[i-2] {
            if check(dp[i-3]) {
                dp[i][1] = true
            }
        }

        if nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1 {
            if check(dp[i-3]) {
                dp[i][1] = true
            }
        }
    }

    return check(dp[len(nums)-1])
}

