func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := []int{}

	for x > 0 {
		str = append(str, x%10)
		x /= 10
	}

	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

