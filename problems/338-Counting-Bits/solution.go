func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		num := i
		for num > 0 {
			if num&0b1 == 1 {
				result[i]++
			}
			num >>= 1
		}
	}
	return result
}

