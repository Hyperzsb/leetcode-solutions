func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(a, b int) bool {
		if points[a][0] < points[b][0] {
			return true
		} else if points[a][0] == points[b][0] {
			return points[a][1] < points[b][1]
		} else {
			return false
		}
	})

	result := 0
	for i := 0; i < len(points); {
		end := points[i][1]
		for i < len(points) && points[i][0] <= end {
			if points[i][1] < end {
				end = points[i][1]
			}
			i++
		}
		result++
	}

	return result
}

