func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}

