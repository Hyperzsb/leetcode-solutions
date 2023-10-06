func numIdenticalPairs(nums []int) int {
    cnts := make(map[int]int)
    for i := range nums {
        cnts[nums[i]]++
    }

    result := 0
    for _, cnt := range cnts {
        result += cnt * (cnt - 1) / 2
    }

    return result
}

