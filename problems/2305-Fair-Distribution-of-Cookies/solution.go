func backtrace(idx int, cookies []int, allocation []int, result *int) {
    if idx == len(cookies) {
        max := 0
        for i := range allocation {
            if max < (allocation)[i] {
                max = (allocation)[i]
            }
        }

        if *result > max {
            *result = max
        }

        return
    }

    for i := range allocation {
        allocation[i] += cookies[idx]
        backtrace(idx+1, cookies, allocation, result)
        allocation[i] -= cookies[idx]
    }
}

func distributeCookies(cookies []int, k int) int {
    result, allocation := int(1e6), make([]int, k)
    backtrace(0, cookies, allocation, &result)

    return result
}

