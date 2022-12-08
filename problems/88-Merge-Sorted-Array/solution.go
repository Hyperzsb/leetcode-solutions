func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m)
	copy(nums3, nums1[:m])

	idx, idxa, idxb := 0, 0, 0
	for idxa < m || idxb < n {
		if idxa >= m {
			nums1[idx] = nums2[idxb]
			idxb++
		} else if idxb >= n {
			nums1[idx] = nums3[idxa]
			idxa++
		} else {
			if nums3[idxa] <= nums2[idxb] {
				nums1[idx] = nums3[idxa]
				idxa++
			} else {
				nums1[idx] = nums2[idxb]
				idxb++
			}
		}

		idx++
	}
}

