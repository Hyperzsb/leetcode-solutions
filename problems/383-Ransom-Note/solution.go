func canConstruct(ransomNote string, magazine string) bool {
	charMap := make(map[byte]int)
	for _, char := range magazine {
		if _, exist := charMap[byte(char)]; exist {
			charMap[byte(char)]++
		} else {
			charMap[byte(char)] = 1
		}
	}

	for _, char := range ransomNote {
		if _, exist := charMap[byte(char)]; exist {
			if charMap[byte(char)] < 1 {
				return false
			} else {
				charMap[byte(char)]--
			}
		} else {
			return false
		}
	}

	return true
}

