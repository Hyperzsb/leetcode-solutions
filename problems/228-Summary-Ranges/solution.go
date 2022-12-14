func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	if len(nums) == 0 {
		return result
	}

	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > end+1 {
			if start < end {
				result = append(result, fmt.Sprintf("%d->%d", start, end))
			} else {
				result = append(result, fmt.Sprintf("%d", end))
			}
			start = nums[i]
			end = nums[i]
		} else {
			end = nums[i]
		}
	}

	if start < end {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	} else {
		result = append(result, fmt.Sprintf("%d", end))
	}

	return result
}

