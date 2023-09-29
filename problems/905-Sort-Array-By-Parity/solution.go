func sortArrayByParity(nums []int) []int {
    for {
        oddIdx := 0
        for oddIdx < len(nums) && nums[oddIdx]%2 == 0 {
            oddIdx++
        }

        evenIdx := oddIdx
        for evenIdx < len(nums) && nums[evenIdx]%2 != 0 {
            evenIdx++
        }

        if oddIdx == evenIdx || evenIdx == len(nums) {
            break
        } else {
            nums[oddIdx] ^= nums[evenIdx]
            nums[evenIdx] ^= nums[oddIdx]
            nums[oddIdx] ^= nums[evenIdx]
        }
    }

    return nums
}

