func bs(nums []int, t int) int {
    s, e := 0, len(nums)-1
    for s < e {
        h := (s + e) / 2

        if nums[h] >= t {
            e = h
        } else {
            s = h + 1
        }
    }

    if nums[s] == t {
        return s
    } else {
        return -1
    }
}

func backtrace(idx, step int, stones []int, dp []map[int]bool) bool {
    if idx >= len(stones) {
        return false
    }

    if idx == len(stones)-1 {
        return true
    }

    if val, ok := dp[idx][step]; ok {
        return val
    }

    result := false
    if step-1 > 0 {
        if next := bs(stones, stones[idx]+step-1); next != -1 {
            result = result || backtrace(next, step-1, stones, dp)
        }
    }

    if !result && step > 0 {
        if next := bs(stones, stones[idx]+step); next != -1 {
            result = result || backtrace(next, step, stones, dp)
        }
    }

    if !result {
        if next := bs(stones, stones[idx]+step+1); next != -1 {
            result = result || backtrace(next, step+1, stones, dp)
        }
    }

    dp[idx][step] = result

    return result
}

func canCross(stones []int) bool {
    dp := make([]map[int]bool, len(stones))
    for i := range dp {
        dp[i] = make(map[int]bool)
    }

    return backtrace(0, 0, stones, dp)
}

