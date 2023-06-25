func max(nums []int) int {
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if max < nums[i] {
            max = nums[i]
        }
    }

    return max
}

func lowerBound(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    start, end := 0, len(nums)-1

    for start < end {
        half := (start + end) / 2

        if target <= nums[half] {
            end = half
        } else {
            start = half + 1
        }
    }

    if nums[start] >= target {
        return start - 1
    } else {
        return start
    }
}

func longestArithSeqLength(nums []int) int {
    hm := make(map[int][]int)
    dp := make([][]int, len(nums))
    maxNum := max(nums)
    for i := range dp {
        if hm[nums[i]] == nil {
            hm[nums[i]] = make([]int, 1)
            hm[nums[i]][0] = i
        } else {
            hm[nums[i]] = append(hm[nums[i]], i)
        }

        dp[i] = make([]int, 2*maxNum+1)
        for j := range dp[i] {
            dp[i][j] = 1
        }
    }

    result := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j <= maxNum && nums[i]-j >= 0; j++ {
            idxes := hm[nums[i]-j]
            lb := lowerBound(idxes, i)

            if lb != -1 {
                dp[i][j] = dp[idxes[lb]][j]+1
            }

            if result < dp[i][j] {
                result = dp[i][j]
            }
        }

        for j := 1; j <= maxNum && nums[i]+j <= maxNum; j++ {
            idxes := hm[nums[i]+j]
            lb := lowerBound(idxes, i)

            if lb != -1 {
                dp[i][maxNum+j] = dp[idxes[lb]][maxNum+j]+1
            }

            if result < dp[i][maxNum+j] {
                result = dp[i][maxNum+j]
            }
        }
    }

    return result
}

