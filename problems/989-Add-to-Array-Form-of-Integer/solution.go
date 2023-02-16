func addToArrayForm(num []int, k int) []int {
	sum, cum := 0, 0

	for i := len(num) - 1; i >= 0; i-- {
		sum = (num[i] + k%10 + cum) % 10
		cum = (num[i] + k%10 + cum) / 10
		num[i] = sum
		k /= 10
	}

	if k > 0 && cum > 0 {
		remain := make([]int, 0)
		for k > 0 {
			remain = append(remain, k%10)
			k /= 10
		}

		for i := 0; i < len(remain)/2; i++ {
			remain[i] ^= remain[len(remain)-1-i]
			remain[len(remain)-1-i] ^= remain[i]
			remain[i] ^= remain[len(remain)-1-i]
		}

		remain = addToArrayForm(remain, cum)

		return append(remain, num...)
	} else if k == 0 && cum > 0 {
		num = append(make([]int, 1), num...)
		num[0] = cum

		return num
	} else if k > 0 && cum == 0 {
		remain := make([]int, 0)
		for k > 0 {
			remain = append(remain, k%10)
			k /= 10
		}

		for i := 0; i < len(remain)/2; i++ {
			remain[i] ^= remain[len(remain)-1-i]
			remain[len(remain)-1-i] ^= remain[i]
			remain[i] ^= remain[len(remain)-1-i]
		}

		return append(remain, num...)
	} else {
		return num
	}
}

