func firstMissingPositive(nums []int) int {
    for i := 0; i < len(nums); {
        idx := nums[i] - 1
        if nums[i] <= len(nums) && nums[i] > 0 && nums[i] != nums[idx] {
            nums[i] ^= nums[idx]
            nums[idx] ^= nums[i]
            nums[i] ^= nums[idx]
        } else {
            i++
        }
    }

    for i := 0; i < len(nums); i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }

    return len(nums)+1
}

