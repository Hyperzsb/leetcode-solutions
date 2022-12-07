func bs(nums []int, target int, offset int) int {
	length := len(nums)

	start, end, mid := 0, length-1, (length-1)/2
	if nums[0] <= nums[length-1] {
		for start < end {
			mid := (start + end) / 2
			if nums[mid] == target {
				return offset + mid
			} else if nums[mid] > target {
				end = mid
			} else {
				start = mid + 1
			}
		}

		mid := (start + end) / 2
		if nums[mid] == target {
			return offset + mid
		} else {
			return -1
		}
	} else {
		result := bs(nums[:mid+1], target, offset)
		if result == -1 {
			return bs(nums[mid+1:], target, offset+mid+1)
		} else {
			return result
		}
	}
}

func search(nums []int, target int) int {
	return bs(nums, target, 0)
}

