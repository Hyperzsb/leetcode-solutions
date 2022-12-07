func addBinary(a string, b string) string {
	if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}

	idxa, idxb := len(a)-1, len(b)-1
	var increase, sum byte = 0, 0
	var result string = ""

	for idxb >= 0 {
		sum = (a[idxa] - '0') + (b[idxb] - '0') + increase
		result = string(sum%2+'0') + result
		increase = sum / 2
		idxa--
		idxb--
	}

	for idxa >= 0 {
		sum = (a[idxa] - '0') + increase
		result = string(sum%2+'0') + result
		increase = sum / 2
		idxa--
	}

	if increase > 0 {
		result = "1" + result
	}

	return result
}

