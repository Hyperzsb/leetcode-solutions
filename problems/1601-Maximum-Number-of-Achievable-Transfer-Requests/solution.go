func backtrace(idx int, requests [][]int, netChange []int, curr int, result *int) {
    if idx == len(requests) {
        for i := range netChange {
            if netChange[i] != 0 {
                return
            }
        }

        if curr > *result {
            *result = curr
        }

        return
    }

    netChange[requests[idx][0]]--
    netChange[requests[idx][1]]++
    backtrace(idx+1, requests, netChange, curr+1, result)
    netChange[requests[idx][0]]++
    netChange[requests[idx][1]]--

    backtrace(idx+1, requests, netChange, curr, result)
}

func maximumRequests(n int, requests [][]int) int {
    result, netChange := 0, make([]int, n)
    backtrace(0, requests, netChange, 0, &result)

    return result
}

