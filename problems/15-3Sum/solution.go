func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    result, resultMap := make([][]int, 0), make(map[[3]int]bool)
    for i := 0; i < len(nums) && nums[i] <= 0 ; i++ {
        for j := i + 1; j < len(nums) && nums[i]+nums[j] <= 0; j++ {
            for k := j + 1; k < len(nums) && nums[i]+nums[j]+nums[k] <= 0; k++ {
                if nums[i]+nums[j]+nums[k] == 0 && !resultMap[[3]int{nums[i], nums[j], nums[k]}] {
                    result = append(result, []int{nums[i], nums[j], nums[k]})
                    resultMap[[3]int{nums[i], nums[j], nums[k]}] = true
                }
            }

            jIdx := j
            for jIdx < len(nums) && nums[jIdx] == nums[j] {
                jIdx++
            }

            if jIdx-1 > i {
                j = jIdx - 1
            }
        }

        iIdx := i
        for iIdx < len(nums) && nums[iIdx] == nums[i] {
            iIdx++
        }

        if iIdx-2 > i {
            i = iIdx - 2
        }
    }

    return result
}

