func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	start, mid, end := 1, (1+num)/2, num
	for start < end {
		mid = (start + end) / 2
		if num == mid*mid {
			return true
		} else if num < mid*mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return false
}

