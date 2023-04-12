func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	start, end, preMin, preMax := -1, 0, -1, -1

	var result int64
	for end < len(nums) {
		if nums[end] < minK || nums[end] > maxK {
			start = end
			end++
			continue
		}

		if nums[end] == minK {
			preMin = end
		}

		if nums[end] == maxK {
			preMax = end
		}

		result += int64(max(0, min(preMin, preMax)-start))
		end++
	}

	return result
}

