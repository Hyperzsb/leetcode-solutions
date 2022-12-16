func area(left, right int, height *[]int) int {
	if (*height)[left] <= (*height)[right] {
		return (right - left) * (*height)[left]
	} else {
		return (right - left) * (*height)[right]
	}

}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := area(left, right, &height)

	for left < right {
		if area(left, right, &height) > result {
			result = area(left, right, &height)
		}

		minHeight := 0
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		for i := left; i <= right; i++ {
			left = i
			if height[i] > minHeight {
				break
			}
		}

		for i := right; i >= left; i-- {
			right = i
			if height[i] > minHeight {
				break
			}
		}
	}

	return result
}

