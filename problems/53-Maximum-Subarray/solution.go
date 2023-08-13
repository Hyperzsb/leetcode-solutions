func maxSubArray(nums []int) int {
    result, sum := nums[0], 0

    for i := 0; i < len(nums); i++ {
        sum += nums[i]

        if sum > result {
            result = sum
        }

        if sum < 0 {
            sum = 0
        }
    }

    return result
}

