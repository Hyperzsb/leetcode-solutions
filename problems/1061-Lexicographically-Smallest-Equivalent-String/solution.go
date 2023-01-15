func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	charEq := make(map[byte]map[byte]bool, 26)
	for i := 0; i < 26; i++ {
		charEq[byte('a'+i)] = make(map[byte]bool)
		charEq[byte('a'+i)][byte('a'+i)] = true
	}

	for i := 0; i < len(s1); i++ {
		char1, char2 := s1[i], s2[i]
		charEq[char1][char2] = true
		charEq[char2][char1] = true

		for char1c, _ := range charEq[char1] {
			for char2c, _ := range charEq[char2] {
				charEq[char1c][char2c] = true
			}
		}

		for char2c, _ := range charEq[char2] {
			for char1c, _ := range charEq[char1] {
				charEq[char2c][char1c] = true
			}
		}
	}

	charSub := make(map[byte]byte, 26)
	for i := 0; i < 26; i++ {
		for j := 0; j <= i; j++ {
			if charEq[byte('a'+i)][byte('a'+j)] {
				charSub[byte('a'+i)] = byte('a' + j)
				break
			}
		}
	}

	result := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		result[i] = charSub[baseStr[i]]
	}

	return string(result)
}

