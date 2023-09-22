func isSubsequence(s string, t string) bool {
    sIdx, tIdx := 0, 0
    for sIdx < len(s) && tIdx < len(t) {
        for tIdx < len(t) {
            if s[sIdx] != t[tIdx] {
                tIdx++
            } else {
                tIdx++
                sIdx++
                break
            }
        }
    }

    if sIdx < len(s) {
        return false
    } else {
        return true
    }
}

