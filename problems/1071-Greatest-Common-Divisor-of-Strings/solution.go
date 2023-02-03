func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isDivisor(target, divisor string) bool {
	if len(target)%len(divisor) != 0 {
		return false
	}

	for i := 0; i < len(target)/len(divisor); i++ {
		if string(target[len(divisor)*i:len(divisor)*(i+1)]) != divisor {
			return false
		}
	}

	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	result := ""
	for i := min(len(str1), len(str2)); i > 0; i-- {
		if len(str1)%i == 0 && len(str2)%i == 0 {
			if isDivisor(str1, string(str1[:i])) && isDivisor(str2, string(str1[:i])) {
				result = string(str1[:i])
				break
			}
		}
	}

	return result
}

