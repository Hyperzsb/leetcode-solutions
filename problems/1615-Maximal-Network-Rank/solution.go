func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maximalNetworkRank(n int, roads [][]int) int {
    isConnected := make(map[[2]int]bool)
    nodes := make([]int, n)
    for _, road := range roads {
        nodes[road[0]]++
        nodes[road[1]]++

        if road[0] < road[1] {
            isConnected[[2]int{road[0], road[1]}] = true
        } else {
            isConnected[[2]int{road[1], road[0]}] = true
        }
    }

    result := 0
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isConnected[[2]int{i, j}] {
                result = max(result, nodes[i]+nodes[j]-1)
            } else {
                result = max(result, nodes[i]+nodes[j])
            }
        }
    }

    return result
}

