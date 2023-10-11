func bs(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }

    s, e := 0, len(nums)-1
    for s < e {
        h := (s + e) / 2

        if nums[h] >= target {
            e = h
        } else {
            s = h + 1
        }
    }

    if nums[s] >= target {
        return s
    } else {
        return s + 1
    }
}

func searchRange(nums []int, target int) []int {
    s, e := bs(nums, target), bs(nums, target+1)

    if s == len(nums) || nums[s] != target {
        return []int{-1, -1}
    } else {
        return []int{s, e-1}
    }
}

