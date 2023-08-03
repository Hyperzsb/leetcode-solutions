func isValid(n int, batteries []int, target int64) bool {
    sum := int64(0)
    for i := range batteries {
        if int64(batteries[i]) >= target {
            sum += target
        } else {
            sum += int64(batteries[i])
        }
    }

    return sum >= int64(n) * target
}

func maxRunTime(n int, batteries []int) int64 {
    sum := int64(0)
    for i := range batteries {
        sum += int64(batteries[i])
    }

    start, end := int64(0), sum / int64(n)
    for start < end {
        half := (start + end) / 2

        if !isValid(n, batteries, half) {
            end = half
        } else {
            start = half + 1
        }
    }

    if isValid(n, batteries, start) {
        return start
    } else {
        return start - 1
    }
}

