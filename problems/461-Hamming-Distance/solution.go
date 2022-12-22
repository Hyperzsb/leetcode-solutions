func hammingDistance(x int, y int) int {
	result := 0
	for x > 0 || y > 0 {
		if (x&0b1)^(y&0b1) == 1 {
			result++
		}
		x >>= 1
		y >>= 1
	}

	return result
}

