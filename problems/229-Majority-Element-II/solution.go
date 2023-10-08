func majorityElement(nums []int) []int {
    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    result := make([]int, 0)
    s, e := 0, 1
    for e < len(nums) {
        if nums[s] != nums[e] {
            if e-s > len(nums)/3 {
                result = append(result, nums[s])
            }
            
            s = e
        }

        e++
    }

    if e-s > len(nums)/3 {
        result = append(result, nums[s])
    }

    return result
}

