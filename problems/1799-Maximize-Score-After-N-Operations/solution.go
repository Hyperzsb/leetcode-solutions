func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func backtrace(idx int, nums []int, numMask []byte, resultMap map[string]int) int {
	if result, ok := resultMap[string(numMask)]; ok {
		return result
	}

	availNums := make([][]int, 0)
	for i := range numMask {
		if numMask[i] == '1' {
			availNums = append(availNums, []int{i, nums[i]})
		}
	}

	if len(availNums) == 2 {
		result := idx * gcd(availNums[0][1], availNums[1][1])
		resultMap[string(numMask)] = result
		return result
	}

	result := 0

	for i := 0; i < len(availNums); i++ {
		for j := i + 1; j < len(availNums); j++ {
			numMask[availNums[i][0]] = '0'
			numMask[availNums[j][0]] = '0'

			curr := idx * gcd(availNums[i][1], availNums[j][1])
			remain := backtrace(idx+1, nums, numMask, resultMap)

			if result < curr+remain {
				result = curr + remain
			}

			numMask[availNums[i][0]] = '1'
			numMask[availNums[j][0]] = '1'
		}
	}

	resultMap[string(numMask)] = result
	return result
}

func maxScore(nums []int) int {
	numMask := make([]byte, len(nums))
	for i := range numMask {
		numMask[i] = '1'
	}

	resultMap := make(map[string]int)

	return backtrace(1, nums, numMask, resultMap)
}

