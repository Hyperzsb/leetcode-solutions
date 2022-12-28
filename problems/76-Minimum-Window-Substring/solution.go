func check(a, b *map[byte]int) bool {
	for key, valb := range *b {
		if vala, exist := (*a)[key]; exist {
			if vala < valb {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	sChar, tChar := make(map[byte]int, 52), make(map[byte]int, 52)
	for i := 0; i < len(t); i++ {
		if _, exist := sChar[s[i]]; exist {
			sChar[s[i]]++
		} else {
			sChar[s[i]] = 1
		}
		if _, exist := tChar[t[i]]; exist {
			tChar[t[i]]++
		} else {
			tChar[t[i]] = 1
		}
	}

	result := s + "invalid"
	start, end := 0, len(t)-1
	for end < len(s) {
		if check(&sChar, &tChar) {
			if end-start+1 < len(result) {
				result = string(s[start : end+1])
			}
			sChar[s[start]]--
			start++
		} else {
			end++
			if end < len(s) {
				if _, exist := sChar[s[end]]; exist {
					sChar[s[end]]++
				} else {
					sChar[s[end]] = 1
				}
			}
		}
	}

	if result == s+"invalid" {
		return ""
	}

	return result
}

