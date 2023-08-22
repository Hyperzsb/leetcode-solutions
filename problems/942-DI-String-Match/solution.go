func backtrace(idx int, s string, nums []bool, result *[]int) bool {
    if idx == len(s) {
        for i := range nums {
            if !nums[i] {
                *result = append(*result, i)
                break
            }
        }

        return true
    }

    if s[idx] == 'I' {
        for i := 0; i < len(nums); i++ {
            if !nums[i] {
                nums[i] = true
                *result = append(*result, i)

                if backtrace(idx+1, s, nums, result) {
                    return true
                }

                nums[i] = false
                *result = (*result)[:len(*result)-1]
            }
        }
    } else {
        for i := len(nums)-1; i >= 0; i-- {
            if !nums[i] {
                nums[i] = true
                *result = append(*result, i)

                if backtrace(idx+1, s, nums, result) {
                    return true
                }

                nums[i] = false
                *result = (*result)[:len(*result)-1]
            }
        }
    }

    return false
}

func diStringMatch(s string) []int {
    nums, result := make([]bool, len(s)+1), make([]int, 0)
    backtrace(0, s, nums, &result)

    return result
}

