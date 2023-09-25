func longestStrChain(words []string) int {
    sort.Slice(words, func(a, b int) bool {
        return len(words[a]) < len(words[b])
    })

    dp := make(map[string]int)
    for i := range words {
        dp[words[i]] = 1
    }

    result := 0
    for i := range words {
        length := dp[words[i]]
        for j := 0; j < len(words[i]); j++ {
            key := string(words[i][:j]) + string(words[i][j+1:])
            if val, ok := dp[key]; ok {
                if length < val+1 {
                    length = val + 1
                }
            }
        }

        dp[words[i]] = length

        if result < length {
            result = length
        }
    }

    return result
}

