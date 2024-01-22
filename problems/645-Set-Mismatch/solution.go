func findErrorNums(nums []int) []int {
    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    duplicate, missing := 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            duplicate = nums[i]
        } else if nums[i] == nums[i-1]+2 {
            missing = nums[i] - 1
        } else {
            continue
        }
    } 

    if nums[0] != 1 {
        return []int{duplicate, 1}
    } else if nums[len(nums)-1] != len(nums) {
        return []int{duplicate, len(nums)}
    } else {
        return []int{duplicate, missing}
    }
}

