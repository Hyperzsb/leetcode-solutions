func convertToBase7(num int) string {
	result := make([]byte, 0)

	var remain int
	if num >= 0 {
		remain = num
	} else {
		remain = -num
	}

	for {
		result = append(result, byte(remain%7+'0'))
		remain /= 7

		if remain == 0 {
			break
		}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i] ^= result[len(result)-1-i]
		result[len(result)-1-i] ^= result[i]
		result[i] ^= result[len(result)-1-i]
	}

	if num >= 0 {
		return string(result)
	} else {
		return "-" + string(result)
	}
}

