func romanToInt(s string) int {
    result := 0

    idx := len(s)-1
    for idx >= 0 {
        if s[idx] == 'I' {
            result += 1
        } else {
            break
        }

        idx--
    }

    if idx >= 0 && s[idx] == 'V' {
        if idx > 0 && s[idx-1] == 'I' {
            result += 4
            idx--
        } else {
            result += 5
        }

        idx--
    }

    for idx >= 0 {
        if s[idx] == 'X' {
            if idx > 0 && s[idx-1] == 'I' {
                result += 9
                idx--
            } else {
                result += 10
            }
        } else {
            break
        }

        idx--
    }

    if idx >= 0 && s[idx] == 'L' {
        if idx > 0 && s[idx-1] == 'X' {
            result += 40
            idx--
        } else {
            result += 50
        }
        idx--
    }

    for idx >= 0 {
        if s[idx] == 'C' {
            if idx > 0 && s[idx-1] == 'X' {
                result += 90
                idx--
            } else {
                result += 100
            }
        } else {
            break
        }
        idx--
    }

    if idx >= 0 && s[idx] == 'D' {
        if idx > 0 && s[idx-1] == 'C' {
            result += 400
            idx--
        } else {
            result += 500
        }
        idx--
    }

    for idx >= 0 {
        if s[idx] == 'M' {
            if idx > 0 && s[idx-1] == 'C' {
                result += 900
                idx--
            } else {
                result += 1000
            }
        } else {
            break
        }
        idx--
    }

    return result
}

