func intersection(nums1 []int, nums2 []int) []int {
	numMap, numMapAll := make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		numMap[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if _, exist := numMap[nums2[i]]; exist {
			numMapAll[nums2[i]] = true
		}
	}

	result := make([]int, 0)
	for key, _ := range numMapAll {
		result = append(result, key)
	}

	return result
}

