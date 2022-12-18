func intersect(nums1 []int, nums2 []int) []int {
	numMap1, numMap2 := make(map[int]int), make(map[int]int)
	for _, val := range nums1 {
		if _, exist := numMap1[val]; exist {
			numMap1[val]++
		} else {
			numMap1[val] = 1
		}
	}

	for _, val := range nums2 {
		if _, exist := numMap2[val]; exist {
			numMap2[val]++
		} else {
			numMap2[val] = 1
		}
	}

	result := make([]int, 0)
	if len(numMap1) < len(numMap2) {
		for key, val1 := range numMap1 {
			if val2, exist := numMap2[key]; exist {
				if val1 < val2 {
					for i := 0; i < val1; i++ {
						result = append(result, key)
					}
				} else {
					for i := 0; i < val2; i++ {
						result = append(result, key)
					}
				}
			}
		}
	} else {
		for key, val2 := range numMap2 {
			if val1, exist := numMap1[key]; exist {
				if val1 < val2 {
					for i := 0; i < val1; i++ {
						result = append(result, key)
					}
				} else {
					for i := 0; i < val2; i++ {
						result = append(result, key)
					}
				}
			}
		}
	}

	return result
}

