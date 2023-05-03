func arraySign(nums []int) int {
	negativeCnt := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			negativeCnt++
		}
	}

	if negativeCnt%2 == 0 {
		return 1
	} else {
		return -1
	}
}

