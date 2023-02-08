func equal(s, t *[]int) bool {
	for i := 0; i < 26; i++ {
		if (*s)[i] != (*t)[i] {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {
	sm, tm := make([]int, 26), make([]int, 26)

	for _, char := range s1 {
		sm[char-'a']++
	}

	for start, end := 0, 0; end < len(s2); end++ {
		if sm[s2[end]-'a'] == 0 {
			for i := start; i < end; i++ {
				tm[s2[i]-'a']--
			}

			start = end + 1
			continue
		}

		if tm[s2[end]-'a'] < sm[s2[end]-'a'] {
			tm[s2[end]-'a']++

			if tm[s2[end]-'a'] == sm[s2[end]-'a'] && equal(&sm, &tm) {
				return true
			}

			continue
		}

		for i := start; i < end; i++ {
			if s2[i]-'a' == s2[end]-'a' {
				start = i + 1
				break
			}

			tm[s2[i]-'a']--
		}
	}

	return equal(&sm, &tm)
}

