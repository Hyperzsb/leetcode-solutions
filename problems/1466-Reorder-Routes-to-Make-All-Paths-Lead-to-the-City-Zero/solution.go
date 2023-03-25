type City struct {
	Visited bool
	Roads   []Road
}

type Road struct {
	Dist int
	Dir  int
}

func traverse(idx int, cities []City, result *int) {
	cities[idx].Visited = true

	for _, r := range cities[idx].Roads {
		if cities[r.Dist].Visited {
			continue
		}

		if r.Dir == 1 {
			*result++
		}

		traverse(r.Dist, cities, result)
	}
}

func minReorder(n int, connections [][]int) int {
	cities := make([]City, n)
	for i := 0; i < n; i++ {
		cities[i] = City{false, make([]Road, 0)}
	}

	for _, conn := range connections {
		cities[conn[0]].Roads = append(cities[conn[0]].Roads, Road{conn[1], 1})
		cities[conn[1]].Roads = append(cities[conn[1]].Roads, Road{conn[0], -1})
	}

	result := 0
	traverse(0, cities, &result)

	return result
}

