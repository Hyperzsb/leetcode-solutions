func reverseBits(num uint32) uint32 {
	var result uint32 = 0b00000000000000000000000000000000
	var flag uint32 = 0b00000000000000000000000000000001

	for i := 0; i < 32; i++ {
		var temp uint32 = num >> i
		result |= flag & temp
		if i < 31 {
			result <<= 1
		}
	}

	return result
}

