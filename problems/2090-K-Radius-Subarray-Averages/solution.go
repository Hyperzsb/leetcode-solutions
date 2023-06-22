func getAverages(nums []int, k int) []int {
    result := make([]int, len(nums))
    for i := range result {
        result[i] = -1
    }

    if len(nums) < 2*k+1 {
        return result
    }
    
    sum := 0
    for i := 0; i < 2*k+1; i++ {
        sum += nums[i]
    }

    for i := 2*k; i < len(nums); i++ {
        result[i-k] = sum / (2 * k + 1)

        if i < len(nums)-1 {
            sum += nums[i+1]
            sum -= nums[i-2*k]
        }
    }

    return result
}

