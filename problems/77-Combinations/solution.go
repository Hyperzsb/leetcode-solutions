func backtrace(idx, n, k int, curr []int, result *[][]int) {
    if len(curr) == k-1 {
        combination := make([]int, len(curr))
        combination = append(combination, idx)
        copy(combination, curr)
        *result = append(*result, combination)

        return
    }

    curr = append(curr, idx)
    for i := idx+1; i <= n; i++ {
        backtrace(i, n, k, curr, result)
    }
    curr = curr[:len(curr)-1]
}

func combine(n int, k int) [][]int {
    result, curr := make([][]int, 0), make([]int, 0)
    for i := 1; i <= n; i++ {
        backtrace(i, n, k, curr, &result)
    }

    return result
}

