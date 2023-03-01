func shuffle(nums []int, n int) []int {
	max := 10000

	first, second := 0, n
	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			nums[i] = (nums[first]%max)*max + nums[i]
			first++
		} else {
			nums[i] = (nums[second]%max)*max + nums[i]
			second++
		}
	}

	for i := 0; i < n*2; i++ {
		nums[i] /= max
	}

	return nums
}

