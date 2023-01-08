func maximumCount(nums []int) int {
	negCount, posCount := 0, 0

	for negCount <= len(nums)-1-posCount && (nums[negCount] != 0 || nums[len(nums)-1-posCount] != 0) {
		fmt.Println(negCount, posCount)
		if nums[negCount] < 0 {
			negCount++
		}
		if nums[len(nums)-1-posCount] > 0 {
			posCount++
		}
	}

	if negCount > posCount {
		return negCount
	} else {
		return posCount
	}
}

