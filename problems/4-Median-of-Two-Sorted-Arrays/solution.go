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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	if len1 == 0 {
		if len2%2 == 0 {
			return float64(nums2[(len2-1)/2]+nums2[len2/2]) / 2
		} else {
			return float64(nums2[(len2-1)/2])
		}
	}

	const MIN int = -10000000
	const MAX int = 10000000
	var median float64
	start, end := 0, len1*2-1
	for {
		mid1 := (start + end) / 2
		mid2 := (len1 + len2 - 1 - mid1)

		left1, right1 := MIN, MAX
		if mid1 > 0 {
			left1 = nums1[(mid1-1)/2]
		}
		if mid1 < len1*2-1 {
			right1 = nums1[(mid1+1)/2]
		}

		left2, right2 := MIN, MAX
		if mid2 > 0 {
			left2 = nums2[(mid2-1)/2]
		}
		if mid2 < len2*2-1 {
			right2 = nums2[(mid2+1)/2]
		}

		fmt.Println(left1, right1, left2, right2)

		if left1 > right2 {
			end = mid1
		} else if left2 > right1 {
			start = mid1 + 1
		} else {
			left := max(left1, left2)
			right := min(right1, right2)

			if nums1[mid1/2] < nums2[mid2/2] {
				median = float64(max(left, nums1[mid1/2])+min(right, nums2[mid2/2])) / 2
			} else {
				median = float64(max(left, nums2[mid2/2])+min(right, nums1[mid1/2])) / 2
			}
			break
		}
	}

	return median
}

