func mySqrt(x int) int {
	result := 0
	for result*result <= x {
		result++
	}

	return result - 1
}

