func isValid(s, e int, nums []int, target int) bool {
    sum := 0
    for i := s; i <= e; i++ {
        sum += nums[i]
    }

    return sum >= target
}

func minSubArrayLen(target int, nums []int) int {
    if !isValid(0, len(nums)-1, nums, target) {
        return 0
    }

    start, end, sum := 0, 0, 0
    for i := range nums {
        sum += nums[i]
        
        if sum >= target {
            end = i
            break
        }
    }

    result := end - start + 1

    for {
        sum -= nums[start]
        start++

        for sum < target && end+1 < len(nums) {
            sum += nums[end+1]
            end++
        }

        if sum < target {
            break
        } else {
            if end-start+1 < result {
                result = end - start + 1
            }
        }
    }

    return result
}

