func exists(nums []int, target int) bool {
    s, e := 0, len(nums)-1
    for s < e {
        h := (s + e) / 2

        if target < nums[h] {
            e = h
        } else if target > nums[h] {
            s = h + 1
        } else {
            return true
        }
    }

    return nums[s] == target
}

func search(nums []int, target int) bool {
    if len(nums) == 1 {
        return nums[0] == target
    }

    prev, curr := 0, 1
    for curr < len(nums) {
        if nums[prev] > nums[curr] {
            break
        }

        prev++
        curr++
    }

    fmt.Println(prev, curr)

    if curr == len(nums) {
        if nums[0] == nums[len(nums)-1] {
            return nums[0] == target
        } else {
            return exists(nums, target)
        }
    }

    return exists(nums[:curr], target) || exists(nums[curr:], target)
}

