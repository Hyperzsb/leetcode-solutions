func reorganizeString(s string) string {
    letterCnt := make([][]int, 26)
    for i := range letterCnt {
        letterCnt[i] = []int{i, 0}
    }

    for i := range s {
        letterCnt[s[i]-'a'][1]++
    }

    sort.Slice(letterCnt, func(a, b int) bool {
        return letterCnt[a][1] > letterCnt[b][1]
    })

    result := make([]byte, 0)
    for letterCnt[0][1] > 0 {
        for letterCnt[0][1] > 0 {
            result = append(result, byte(letterCnt[0][0])+'a')
            letterCnt[0][1]--
            
            for i := 1; i < 26 && letterCnt[i][1] > 0 && letterCnt[0][1] > 0; i++ {
                cnt := letterCnt[i][1]
                for j := 0; j < cnt && letterCnt[0][1] > 0; j++ {
                    result = append(result, byte(letterCnt[i][0])+'a')
                    letterCnt[i][1]--
                    result = append(result, byte(letterCnt[0][0])+'a')
                    letterCnt[0][1]--
                }
            }


            if letterCnt[0][1] > 0 {
                return ""
            }
        }

        sort.Slice(letterCnt, func(a, b int) bool {
            return letterCnt[a][1] > letterCnt[b][1]
        })
    }

    return string(result)
}

