func backtrace(s string, wordMap map[string]bool, dp []int) bool {
    if dp[len(s)] != 0 {
        if dp[len(s)] == 1 {
            return true
        } else {
            return false
        }
    }

    for i := 0; i < len(s); i++ {
        if wordMap[string(s[:i+1])] {
            if backtrace(string(s[i+1:]), wordMap, dp) {
                dp[len(s)] = 1

                return true
            }
        }
    }

    dp[len(s)] = -1

    return false
}

func wordBreak(s string, wordDict []string) bool {
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }

    dp := make([]int, len(s)+1)
    dp[0] = 1

    return backtrace(s, wordMap, dp)
}

