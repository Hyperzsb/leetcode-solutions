type Computer struct {
	Visited   bool
	Neighbors []int
}

func traverse(idx int, computers []Computer, computerCnt, cableCnt *int) {
	computers[idx].Visited = true
	*computerCnt += 1
	*cableCnt += len(computers[idx].Neighbors)

	for _, n := range computers[idx].Neighbors {
		if !computers[n].Visited {
			traverse(n, computers, computerCnt, cableCnt)
		}
	}
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	computers := make([]Computer, n)
	for i := 0; i < n; i++ {
		computers[i] = Computer{false, make([]int, 0)}
	}

	for _, conn := range connections {
		computers[conn[0]].Neighbors = append(computers[conn[0]].Neighbors, conn[1])
		computers[conn[1]].Neighbors = append(computers[conn[1]].Neighbors, conn[0])
	}

	visited := 0
	networks, availCables := 0, 0
	for visited < n {
		idx := 0
		for i := 0; i < n; i++ {
			if !computers[i].Visited {
				idx = i
				break
			}
		}

		computerCnt, cableCnt := 0, 0
		traverse(idx, computers, &computerCnt, &cableCnt)
		visited += computerCnt
		cableCnt /= 2

		networks++
		availCables += cableCnt - (computerCnt - 1)
	}

	if networks-1 > availCables {
		return -1
	} else {
		return networks - 1
	}
}

