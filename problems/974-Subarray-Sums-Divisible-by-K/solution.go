func subarraysDivByK(nums []int, k int) int {
	modulusMap := make(map[int]int)
	modulusMap[0] = 1

	result := 0
	modulus := 0
	for i := 0; i < len(nums); i++ {
		modulus = (modulus + (nums[i] % k) + k) % k
		result += modulusMap[modulus]
		modulusMap[modulus]++
	}

	return result
}

