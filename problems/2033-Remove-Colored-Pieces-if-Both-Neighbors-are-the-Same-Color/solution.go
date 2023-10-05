func winnerOfGame(colors string) bool {
    aCnt, bCnt := 0, 0
    for i := 1; i < len(colors)-1; i++ {
        if colors[i-1] == 'A' && colors[i] == 'A' && colors[i+1] == 'A' {
            aCnt++
        }

        if colors[i-1] == 'B' && colors[i] == 'B' && colors[i+1] == 'B' {
            bCnt++
        }
    }

    return aCnt > bCnt
}

