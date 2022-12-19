type Char struct {
	index int
	count int
}

func firstUniqChar(s string) int {
	charMap := make(map[byte]Char)
	for idx, val := range s {
		if _, exist := charMap[byte(val)]; exist {
			charMap[byte(val)] = Char{idx, charMap[byte(val)].count + 1}
		} else {
			charMap[byte(val)] = Char{idx, 1}
		}
	}

	result := -1
	for _, val := range charMap {
		if val.count == 1 {
			if val.index < result || result == -1 {
				result = val.index
			}
		}
	}

	return result
}

