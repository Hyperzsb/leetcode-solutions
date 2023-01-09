func countDigits(num int) int {
	result := 0
	temp := num

	for temp > 0 {
		digit := temp % 10
		if num%digit == 0 {
			result++
		}
		temp /= 10
	}

	return result
}

