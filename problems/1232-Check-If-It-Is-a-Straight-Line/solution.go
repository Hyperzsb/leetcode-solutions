func slope(a, b []int) float64 {
	return float64(b[1]-a[1]) / float64(b[0]-a[0])
}

func checkStraightLine(coordinates [][]int) bool {
	x := coordinates[0][0]
	flag := true
	for i := range coordinates {
		if coordinates[i][0] != x {
			flag = false
			break
		}
	}

	if flag {
		return true
	}

	sort.Slice(coordinates, func(a, b int) bool {
		if coordinates[a][0] < coordinates[b][0] {
			return true
		} else if coordinates[a][0] > coordinates[b][0] {
			return false
		} else {
			return coordinates[a][1] < coordinates[b][1]
		}
	})

	k := slope(coordinates[0], coordinates[1])

	for i := 2; i < len(coordinates); i++ {
		if slope(coordinates[0], coordinates[i]) != k {
			return false
		}
	}

	return true
}

