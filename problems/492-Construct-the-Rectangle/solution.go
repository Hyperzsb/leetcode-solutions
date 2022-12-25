func constructRectangle(area int) []int {
	result := make([]int, 2)
	for i := int(math.Sqrt(float64(area))); i >= 1; i-- {
		if area%i == 0 {
			result[0] = area / i
			result[1] = i
			break
		}
	}

	return result
}

