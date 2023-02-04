func convert(raw string, mapping *map[byte]byte) string {
	result := make([]byte, len(raw))
	for i := 0; i < len(raw); i++ {
		result[i] = (*mapping)[raw[i]]
	}

	return string(result)
}

func isSorted(former, latter string) bool {
	for i := 0; i < len(former) && i < len(latter); i++ {
		if former[i] < latter[i] {
			return true
		} else if former[i] == latter[i] {
			continue
		} else {
			return false
		}
	}

	return len(former) <= len(latter)
}

func isAlienSorted(words []string, order string) bool {
	mapping := make(map[byte]byte)
	for i := 0; i < 26; i++ {
		mapping[order[i]] = byte('a' + i)
	}

	for i := 0; i < len(words); i++ {
		words[i] = convert(words[i], &mapping)
	}

	if len(words) == 1 {
		return true
	}

	for i := 1; i < len(words); i++ {
		if !isSorted(words[i-1], words[i]) {
			return false
		}
	}

	return true
}

