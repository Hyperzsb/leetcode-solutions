func thirdMax(nums []int) int {
	numsQ := make([]int, 1)
	numsQ[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if len(numsQ) == 1 {
			if nums[i] >= numsQ[0] {
				if nums[i] != numsQ[0] {
					numsQ = append(numsQ, numsQ[0])
					numsQ[0] = nums[i]
				}
			} else {
				numsQ = append(numsQ, nums[i])
			}
		} else if len(numsQ) == 2 {
			if nums[i] >= numsQ[0] {
				if nums[i] != numsQ[0] {
					numsQ = append(numsQ, numsQ[1])
					numsQ[1] = numsQ[0]
					numsQ[0] = nums[i]
				}
			} else if nums[i] < numsQ[0] && nums[i] >= numsQ[1] {
				if nums[i] != numsQ[1] {
					numsQ = append(numsQ, numsQ[1])
					numsQ[1] = nums[i]
				}
			} else {
				numsQ = append(numsQ, nums[i])
			}
		} else {
			if nums[i] >= numsQ[0] {
				if nums[i] != numsQ[0] {
					numsQ[2] = numsQ[1]
					numsQ[1] = numsQ[0]
					numsQ[0] = nums[i]
				}
			} else if nums[i] < numsQ[0] && nums[i] >= numsQ[1] {
				if nums[i] != numsQ[1] {
					numsQ[2] = numsQ[1]
					numsQ[1] = nums[i]
				}
			} else if nums[i] < numsQ[1] && nums[i] >= numsQ[2] {
				if nums[i] != numsQ[2] {
					numsQ[2] = nums[i]
				}
			} else {
				continue
			}
		}
	}

	if len(numsQ) < 3 {
		return numsQ[0]
	} else {
		return numsQ[2]
	}
}

