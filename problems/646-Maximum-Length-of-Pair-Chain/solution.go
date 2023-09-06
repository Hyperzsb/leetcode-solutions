func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(a, b int) bool {
        if pairs[a][1] < pairs[b][1] {
            return true
        } else if pairs[a][1] > pairs[b][1] {
            return false
        } else {
            return pairs[a][0] < pairs[b][0] 
        }
    })

    result, right := 1, pairs[0][1]
    for i := 1; i < len(pairs); i++ {
        if pairs[i][0] > right {
            result++
            right = pairs[i][1]
        }
    }

    return result
}

