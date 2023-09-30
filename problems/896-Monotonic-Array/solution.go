func isMonotonic(nums []int) bool {
    mode := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            if mode == -1 {
                return false
            } else {
                mode = 1
            }
        } else if nums[i-1] > nums[i] {
            if mode == 1 {
                return false
            } else {
                mode = -1
            }
        } else {
            continue
        }
    }

    return true
}

