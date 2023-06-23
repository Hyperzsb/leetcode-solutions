type Number struct {
    Val int
    Cost int
}

func calcCost(target int, numbers []Number) int64 {
    result := int64(0)

    for _, number := range numbers {
        if number.Val >= target {
            result += int64((number.Val - target) * number.Cost)
        } else {
            result += int64((target - number.Val) * number.Cost)
        }
    }

    return result
}

func min(nums []int64) int64 {
    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })

    return nums[0]
}

func minCost(nums []int, cost []int) int64 {
    numbers := make([]Number, len(nums))
    for i := range numbers {
        numbers[i].Val = nums[i]
        numbers[i].Cost = cost[i]
    }

    sort.Slice(numbers, func(a, b int) bool {
        return numbers[a].Val < numbers[b].Val
    })

    result := int64(1<<52)
    start, end := numbers[0].Val, numbers[len(nums)-1].Val
    for start <= end {
        half := (start + end) / 2
        l := calcCost(half, numbers)
        g := calcCost(half+1, numbers)

        result = min([]int64{result, l, g})
        if l < g {
            end = half - 1
        } else {
            start = half + 1
        }
    }

    return result
}

