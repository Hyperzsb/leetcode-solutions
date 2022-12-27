func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	start, end, idx := 0, len(nums)-1, len(nums)-1
	for start <= end {
		squareS, squareE := nums[start]*nums[start], nums[end]*nums[end]
		if squareS < squareE {
			result[idx] = squareE
			end--
		} else {
			result[idx] = squareS
			start++
		}
		idx--
	}

	return result
}

