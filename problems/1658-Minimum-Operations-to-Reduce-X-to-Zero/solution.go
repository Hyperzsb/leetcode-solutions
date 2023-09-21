func minOperations(nums []int, x int) int {
    target, n := -x, len(nums)
    for _, num := range nums {
        target += num
    }
    
    if target == 0 {
        return n
    }
    
    maxLen, curSum, left := 0, 0, 0
    
    for right, val := range nums {
        curSum += val
        for left <= right && curSum > target {
            curSum -= nums[left]
            left++
        }
        if curSum == target {
            if right - left + 1 > maxLen {
                maxLen = right - left + 1
            }
        }
    }
    
    if maxLen != 0 {
        return n - maxLen
    }
    return -1
}

