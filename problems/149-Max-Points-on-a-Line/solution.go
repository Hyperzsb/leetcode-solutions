func maxPoints(points [][]int) int {
	result := 1
	if len(points) == 1 {
		return result
	}

	for i := 0; i < len(points); i++ {
		slopeMap := make(map[float64]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			var slope float64
			if points[i][0] == points[j][0] {
				slope = 1000000.0
			} else {
				slope = (float64(points[i][1]) - float64(points[j][1])) /
					(float64(points[i][0]) - float64(points[j][0]))
			}
			slopeMap[slope]++

			if slopeMap[slope]+1 > result {
				result = slopeMap[slope] + 1
			}
		}
	}

	return result
}

