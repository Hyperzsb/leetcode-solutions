func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	currentSum, result := 0, 0
	flag := false
	for i := 0; i < len(nums); i++ {
		currentSum = nums[i]
		start, end := 0, len(nums)-1
		for start < end {
			if start == i {
				start++
				continue
			}
			if end == i {
				end--
				continue
			}
			twoSum := nums[start] + nums[end]
			currentSum += twoSum
			if currentSum < target {
				start++
			} else {
				end--
			}
			if !flag || abs(target-currentSum) < abs(target-result) {
				flag = true
				result = currentSum
			}
			currentSum -= twoSum
		}
	}

	return result
}

