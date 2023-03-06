func constructNext(needle string) []int {
	result := make([]int, len(needle))

	start := -1
	result[0] = start
	for end := 1; end < len(needle); end++ {
		for start >= 0 && needle[start+1] != needle[end] {
			start = result[start]
		}

		if needle[start+1] == needle[end] {
			start++
		}

		result[end] = start
	}

	return result
}

func strStr(haystack string, needle string) int {
	next := constructNext(needle)

	nIdx := -1
	for hIdx := 0; hIdx < len(haystack); hIdx++ {
		for nIdx >= 0 && needle[nIdx+1] != haystack[hIdx] {
			nIdx = next[nIdx]
		}

		if needle[nIdx+1] == haystack[hIdx] {
			nIdx++
		}

		if nIdx == len(needle)-1 {
			return hIdx - len(needle) + 1
		}
	}

	return -1
}

