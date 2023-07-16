func putMarbles(weights []int, k int) int64 {
	var min, max int64 = 0, 0

	sums := make([]int64, len(weights)-1)
	for i := 0; i < len(weights)-1; i++ {
		sums[i] = int64(weights[i]) + int64(weights[i+1])
	}

	sort.Slice(sums, func(a, b int) bool {
		return sums[a] < sums[b]
	})

	for i := 0; i < k-1; i++ {
		min += sums[i]
		max += sums[len(sums)-1-i]
	}

	fmt.Println(max, min)

	return max - min
}

