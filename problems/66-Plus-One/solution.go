func plusOne(digits []int) []int {
	var increase, sum int = 1, 0

	for i := len(digits) - 1; i >= 0; i-- {
		sum = digits[i] + increase
		digits[i] = sum % 10
		increase = sum / 10
		if i == 0 && increase > 0 {
			temp := make([]int, 0)
			temp = append(temp, increase)
			temp = append(temp, digits...)
			digits = temp
		}
	}

	return digits
}

