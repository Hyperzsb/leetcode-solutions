func backspaceCompare(s string, t string) bool {
    sResult := make([]byte, 0)
    for i := range s {
        if s[i] == '#' {
            if len(sResult) > 0 {
                sResult = sResult[:len(sResult)-1]
            }
        } else {
            sResult = append(sResult, s[i])
        }
    }

    tResult := make([]byte, 0)
    for i := range t {
        if t[i] == '#' {
            if len(tResult) > 0 {
                tResult = tResult[:len(tResult)-1]
            }
        } else {
            tResult = append(tResult, t[i])
        }
    }

    return string(sResult) == string(tResult)
}

