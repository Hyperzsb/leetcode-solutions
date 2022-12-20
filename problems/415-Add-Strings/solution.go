func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}

	var sum, cum byte
	idx1, idx2 := len(num1)-1, len(num2)-1
	result := make([]byte, 0)
	for idx2 >= 0 {
		sum = ((num1[idx1] - '0') + (num2[idx2] - '0') + cum) % 10
		cum = ((num1[idx1] - '0') + (num2[idx2] - '0') + cum) / 10
		result = append(result, sum+'0')
		idx1--
		idx2--
	}

	for idx1 >= 0 {
		sum = ((num1[idx1] - '0') + cum) % 10
		cum = ((num1[idx1] - '0') + cum) / 10
		result = append(result, sum+'0')
		idx1--
	}

	if cum > 0 {
		sum = cum
		result = append(result, sum+'0')
	}

	for i := 0; i < len(result)/2; i++ {
		result[i] ^= result[len(result)-1-i]
		result[len(result)-1-i] ^= result[i]
		result[i] ^= result[len(result)-1-i]
	}

	return string(result)
}

