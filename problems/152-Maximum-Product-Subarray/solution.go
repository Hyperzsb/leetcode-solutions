func minMax(a, b, c int) (int, int) {
	arr := make([]int, 3)
	arr[0] = a
	arr[1] = b
	arr[2] = c
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	return arr[0], arr[2]
}

func maxProduct(nums []int) int {
	min, max, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		min, max = minMax(min*nums[i], max*nums[i], nums[i])

		if max > result {
			result = max
		}
	}

	return result
}

