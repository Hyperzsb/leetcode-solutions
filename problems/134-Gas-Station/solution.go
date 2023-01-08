func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}

	sum, result := 0, 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i]
		if sum < 0 {
			result = i + 1
			sum = 0
		}
	}

	return result
}

