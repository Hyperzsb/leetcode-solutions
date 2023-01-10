func calc(s string) int {
	result, flag := 0, false
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return len(s)
		}
		if s[i] == '0' {
			if flag {
				result *= 10
			}
		} else {
			if !flag {
				flag = true
			}
			result *= 10
			result += int(s[i] - '0')
		}
	}

	return result
}

func maximumValue(strs []string) int {
	result := 0
	for _, str := range strs {
		val := calc(str)
		if val > result {
			result = val
		}
	}

	return result
}

