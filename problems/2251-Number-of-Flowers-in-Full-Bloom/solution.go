func lowerBound(nums []int, target int) int {
    s, e := 0, len(nums)-1
    for s < e {
        h := (s + e) / 2
        if target <= nums[h] {
            e = h
        } else {
            s = h + 1
        }
    }

    if nums[s] < target {
        return s + 1
    } else {
        return s
    }
}

func upperBound(nums []int, target int) int {
    s, e := 0, len(nums)-1
    for s < e {
        h := (s + e) / 2
        
        if (s+e)%2 == 1 {
            h += 1
        }

        if target >= nums[h] {
            s = h
        } else {
            e = h - 1
        }
    }

    if nums[e] > target {
        return e
    } else {
        return e + 1
    }
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
    bloom, wither := make([]int, len(flowers)), make([]int, len(flowers))
    for i := range flowers {
        bloom[i], wither[i] = flowers[i][0], flowers[i][1]
    }

    sort.Slice(bloom, func(a, b int) bool {
        return bloom[a] < bloom[b]
    })

    sort.Slice(wither, func(a, b int) bool {
        return wither[a] < wither[b]
    })

    result := make([]int, len(people))
    for i := range people {
        result[i] = upperBound(bloom, people[i]) - lowerBound(wither, people[i])
    }

    return result
}

