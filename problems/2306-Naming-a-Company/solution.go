func calc(a, b *map[string]bool) int64 {
	count := 0
	for key, _ := range *a {
		if (*b)[key] {
			count++
		}
	}

	return int64((len(*a) - count) * (len(*b) - count))
}

func distinctNames(ideas []string) int64 {
	headMap := make(map[byte]map[string]bool)
	for _, word := range ideas {
		var tail string
		if len(word) == 1 {
			tail = "#"
		} else {
			tail = string(word[1:])
		}

		if tailMap, exist := headMap[word[0]]; exist {
			tailMap[tail] = true
		} else {
			headMap[word[0]] = make(map[string]bool)
			headMap[word[0]][tail] = true
		}
	}

	headList := make([]map[string]bool, 0, len(headMap))
	for _, val := range headMap {
		headList = append(headList, val)
	}

	var result int64
	for i := 0; i < len(headList); i++ {
		for j := i + 1; j < len(headList); j++ {
			result += calc(&headList[i], &headList[j])
		}
	}

	return result * 2
}

