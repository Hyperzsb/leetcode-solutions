func isValid(n, index, maxSum, target int) bool {
	sum := 0

	if target <= index+1 {
		sum += (1 + target) * target / 2 + (index + 1 - target)
	} else {
		sum += (target - index + target) * (index + 1) / 2
	}

	if target <= n-index {
		sum += (1 + target) * target / 2 + (n - index - target)
	} else {
		sum += (target - (n - index - 1) + target) * (n - index) / 2
	}

	sum -= target

	if sum <= maxSum {
		return true
	} else {
		return false
	}
}

func maxValue(n int, index int, maxSum int) int {
	start, end := 1, maxSum
	for start < end {
		half := (start + end) / 2

		if !isValid(n, index, maxSum, half) {
			end = half
		} else {
			start = half + 1
		}
	}

	if !isValid(n, index, maxSum, start) {
		return start - 1
	} else {
		return start
	}
}

