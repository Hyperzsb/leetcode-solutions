func pivotInteger(n int) int {
	result := math.Sqrt(float64(n*n+n) / 2)
	if result-float64(int(result)) == 0 {
		return int(result)
	} else {
		return -1
	}
}

