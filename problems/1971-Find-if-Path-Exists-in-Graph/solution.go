func validPath(n int, edges [][]int, source int, destination int) bool {
	nodeEdges := make([][]int, n)
	for i := 0; i < n; i++ {
		nodeEdges[i] = make([]int, 0)
	}

	for _, val := range edges {
		nodeEdges[val[0]] = append(nodeEdges[val[0]], val[1])
		nodeEdges[val[1]] = append(nodeEdges[val[1]], val[0])
	}

	nodesVisited := make([]bool, n)

	bfs := make([]int, 1)
	bfs[0] = source

	for i := 0; i < len(bfs); i++ {
		current := bfs[i]
		if current == destination {
			return true
		}
		nodesVisited[current] = true
		for i := 0; i < len(nodeEdges[current]); i++ {
			if !nodesVisited[nodeEdges[current][i]] {
				bfs = append(bfs, nodeEdges[current][i])
			}
		}
	}

	return false
}

