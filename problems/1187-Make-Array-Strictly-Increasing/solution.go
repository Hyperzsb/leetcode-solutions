func upperBound(nums []int, target int) int {
    start, end := 0, len(nums)-1

    for start < end {
        half := (start + end) / 2

        if target <= nums[half] {
            end = half
        } else {
            start = half + 1
        }
    }

    idx := (start + end) / 2
    if nums[idx] <= target {
        return idx + 1
    } else {
        return idx
    }
}

func calc(idx, prev int, arr1, arr2 []int, dp [][]int) int {
    if idx == len(arr1) {
        return 0
    }

    ubIdx := upperBound(arr2, prev)

    fmt.Println(idx, ubIdx)

    if dp[idx][ubIdx] != -1 {
        return dp[idx][ubIdx]
    }

    if ubIdx == len(arr2) && arr1[idx] <= prev {
        dp[idx][ubIdx] = math.MaxInt32
        return dp[idx][ubIdx]
    }

    pick, notpick := math.MaxInt32, math.MaxInt32
    if ubIdx < len(arr2) {
        pick = 1 + calc(idx+1, arr2[ubIdx], arr1, arr2, dp)
    }

    if arr1[idx] > prev {
        notpick = calc(idx+1, arr1[idx], arr1, arr2, dp)
    }

    if pick < notpick {
        dp[idx][ubIdx] = pick
    } else {
        dp[idx][ubIdx] = notpick
    }

    return dp[idx][ubIdx]
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
    sort.Slice(arr2, func(a, b int) bool {
        return arr2[a] < arr2[b]
    })

    dp := make([][]int, len(arr1))
    for i := range dp {
        dp[i] = make([]int, len(arr2)+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    if result := calc(0, -1, arr1, arr2, dp); result == math.MaxInt32 {
        return -1
    } else {
        return result
    }
}

