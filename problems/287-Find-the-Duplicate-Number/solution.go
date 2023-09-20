func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func findDuplicate(nums []int) int {
    result := 0
    for i := range nums {
        if nums[abs(nums[i])] > 0 {
            nums[abs(nums[i])] = -nums[abs(nums[i])]
        } else {
            result = abs(nums[i])
        }
    }

    for i := range nums {
        if nums[i] < 0 {
            nums[i] = -nums[i]
        }
    }

    return result
}

