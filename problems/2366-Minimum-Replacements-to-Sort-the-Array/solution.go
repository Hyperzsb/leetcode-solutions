func minimumReplacement(nums []int) int64 {
    result, last := int64(0), nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] > last {
            factor := nums[i] / last
            
            if nums[i] % last != 0 {
                factor++
            }

            last = nums[i] / factor
            result += int64(factor - 1)
        } else {
            last = nums[i]
        }
    }

    return result
}

