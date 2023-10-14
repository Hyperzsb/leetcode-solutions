func commonChars(words []string) []string {
    globalCnt := make([]int, 26)
    for i := 0; i < 26; i++ {
        globalCnt[i] = math.MaxInt32
    }

    for i := range words {
        localCnt := make([]int, 26)
        for j := range words[i] {
            localCnt[int(words[i][j]-'a')]++
        }

        for j := 0; j < 26; j++ {
            if globalCnt[j] > localCnt[j] {
                globalCnt[j] = localCnt[j]
            }
        }
    }

    result := make([]string, 0)
    for i := range globalCnt {
        for j := 0; j < globalCnt[i]; j++ {
            result = append(result, string(i+'a'))
        }
    }

    return result
}

