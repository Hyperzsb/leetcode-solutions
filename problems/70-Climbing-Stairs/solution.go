func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	steps := make([]int, n)
	steps[n-1] = 1
	steps[n-2] = 2
	for i := n - 3; i >= 0; i-- {
		steps[i] = steps[i+1] + steps[i+2]
	}

	return steps[0]
}

