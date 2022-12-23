func bitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	n, length := num, 0
	for n > 0 {
		length++
		n >>= 1
	}
	mask := 0
	for i := 0; i < length; i++ {
		mask |= 0b1
		mask <<= 1
	}
	mask >>= 1

	return mask ^ num
}

