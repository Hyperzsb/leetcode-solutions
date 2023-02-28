func isAnagram(s, t *map[byte]int) bool {
	for tKey, tVal := range *t {
		if (*s)[tKey] != tVal {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	rawSet, targetSet := make(map[byte]int), make(map[byte]int)
	for _, char := range p {
		targetSet[byte(char)]++
	}

	result := make([]int, 0)

	start, end := 0, 0
	for end < len(p)-1 {
		rawSet[s[end]]++
		end++
	}

	for end < len(s) {
		rawSet[s[end]]++

		if isAnagram(&rawSet, &targetSet) {
			result = append(result, start)
		}

		rawSet[s[start]]--
		start++
		end++
	}

	return result
}

