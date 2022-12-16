func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)

	start, end, length := 0, 0, 0
	for end < len(s) {
		if _, exist := charMap[byte(s[end])]; exist {
			if end-start > length {
				length = end - start
			}

			for i := start; i < charMap[byte(s[end])]; i++ {
				delete(charMap, byte(s[i]))
			}

			start = charMap[byte(s[end])] + 1
			charMap[byte(s[end])] = end
		} else {
			charMap[byte(s[end])] = end
		}
		end++
	}

	if end-start > length {
		length = end - start
	}

	return length
}

