func abs(n int) int {
    if n >= 0 {
        return n
    } else {
        return -n
    }
}

func traverse(s, e int, locs []int, fuel int, dp map[string]int) int {
    if fuel < 0 {
        return 0
    }

    key := fmt.Sprintf("%d:%d:%d", s, e, fuel)
    if val, ok := dp[key]; ok {
        return val
    }

    result := 0
    if s == e {
        result++
    }

    for i := 0; i < len(locs); i++ {
        if i == s {
            continue
        }

        cnt := traverse(i, e, locs, fuel-abs(locs[s]-locs[i]), dp)
        result = (result + cnt) % (1e9 + 7)
    }

    dp[key] = result

    return result
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
    dp := make(map[string]int)
    return traverse(start, finish, locations, fuel, dp)
}

