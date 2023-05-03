func findDifference(nums1 []int, nums2 []int) [][]int {
	numMap1, numMap2 := make(map[int]bool), make(map[int]bool)

	for _, num := range nums1 {
		numMap1[num] = true
	}

	for _, num := range nums2 {
		numMap2[num] = true
	}

	for num, _ := range numMap1 {
		if numMap2[num] {
			delete(numMap1, num)
			delete(numMap2, num)
		}
	}

	result := make([][]int, 2)

	result[0] = make([]int, 0, len(numMap1))
	for num, _ := range numMap1 {
		result[0] = append(result[0], num)
	}

	result[1] = make([]int, 0, len(numMap2))
	for num, _ := range numMap2 {
		result[1] = append(result[1], num)
	}

	return result
}

