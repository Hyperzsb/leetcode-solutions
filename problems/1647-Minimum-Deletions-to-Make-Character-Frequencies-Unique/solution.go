func minDeletions(s string) int {
    charToFreq := make(map[byte]int)
    for i := range s {
        charToFreq[s[i]]++
    }

    freqToChar := make(map[int]int)
    for _, v := range charToFreq {
        freqToChar[v]++
    }

    result := 0
    for {
        flag := true
        for k, v := range freqToChar {
            if v > 1 {
                freqToChar[k]--
                if k > 1 {
                    freqToChar[k-1]++
                }

                result++

                flag = false
                break
            }
        }

        if flag {
            break
        }
    }

    return result
}

