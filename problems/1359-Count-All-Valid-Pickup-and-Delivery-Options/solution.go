func countOrders(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        curr := (2 * (i - 1) + 1 + 1) * (2 * (i - 1) + 1) / 2
        result = (result * curr) % (1e9 + 7)
    }

    return result
}

