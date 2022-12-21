type Stop struct {
	id    int
	count int
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stops := make(map[int][]int)
	stopsVisited := make(map[int]bool)
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			if _, exist := stops[routes[i][j]]; exist {
				stops[routes[i][j]] = append(stops[routes[i][j]], i)
			} else {
				stops[routes[i][j]] = make([]int, 1)
				stops[routes[i][j]][0] = i
				stopsVisited[routes[i][j]] = false
			}
		}
	}

	stopsAvailable := make([]Stop, 1)
	stopsAvailable[0] = Stop{source, 0}
	stopsVisited[source] = true
	for i := 0; i < len(stopsAvailable); i++ {
		for j := 0; j < len(stops[stopsAvailable[i].id]); j++ {
			for k := 0; k < len(routes[stops[stopsAvailable[i].id][j]]); k++ {
				if routes[stops[stopsAvailable[i].id][j]][k] == target {
					return stopsAvailable[i].count + 1
				}
				if !stopsVisited[routes[stops[stopsAvailable[i].id][j]][k]] {
					stopsAvailable = append(stopsAvailable, Stop{
						routes[stops[stopsAvailable[i].id][j]][k],
						stopsAvailable[i].count + 1})
					stopsVisited[routes[stops[stopsAvailable[i].id][j]][k]] = true
				}
			}
			routes[stops[stopsAvailable[i].id][j]] = make([]int, 0)
		}
	}

	return -1
}

