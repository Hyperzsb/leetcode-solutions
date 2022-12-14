func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 0 {
		if n%2 == 1 && n != 1 {
			return false
		}
		n /= 2
	}

	return true
}

