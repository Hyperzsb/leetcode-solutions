func hammingWeight(num uint32) int {
	var flag uint32 = 0b00000000000000000000000000000001
	result := 0
	for i := 0; i < 32; i++ {
		if flag&num > 0 {
			result++
		}
		flag <<= 1
	}

	return result
}

