func isValid(nums []int, p, target int) bool {
    cnt := 0
    for i := 0; i < len(nums)-1; i++ {
        if nums[i+1]-nums[i] <= target {
            cnt++
            i++
        }

        if cnt >= p {
            return true
        }
    }

    return false
}

func minimizeMax(nums []int, p int) int {
    if len(nums) < 2 {
        return 0
    }

    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    start, end := 0, nums[len(nums)-1] - nums[0]
    for start < end {
        half := (start + end) / 2

        if isValid(nums, p, half) {
            end = half
        } else {
            start = half + 1
        }
    }

    return start
}

