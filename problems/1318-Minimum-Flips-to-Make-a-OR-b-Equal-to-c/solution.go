func minFlips(a int, b int, c int) int {
	result := 0
	for a > 0 || b > 0 || c > 0 {
		aBit := a & 1
		bBit := b & 1
		cBit := c & 1

		if cBit == 1 {
			if aBit == 0 && bBit == 0 {
				result++
			}
		} else {
			if aBit == 1 {
				result++
			}

			if bBit == 1 {
				result++
			}
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return result
}

