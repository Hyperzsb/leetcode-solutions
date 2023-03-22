func getCnt(n int) int64 {
	return int64((n + 1) * n / 2)
}

func zeroFilledSubarray(nums []int) int64 {
	var result int64

	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			result += getCnt(i - start)
			for i < len(nums) && nums[i] != 0 {
				i++
			}
			start = i
		}
	}
	result += getCnt(len(nums) - start)

	return result
}

