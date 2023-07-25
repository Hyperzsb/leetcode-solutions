func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(a, b int) bool {
        if intervals[a][1] < intervals[b][1] {
            return true
        } else if intervals[a][1] == intervals[b][1] {
            return intervals[a][0] < intervals[b][0]
        } else {
            return false
        }
    })

    result := 0

    prev := 0
    for i := 1; i < len(intervals); i++ {
        if intervals[prev][1] > intervals[i][0] {
            result++
        } else {
            prev = i
        }
    }

    return result
}

