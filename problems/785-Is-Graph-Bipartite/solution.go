type MNode struct {
	Color     int
	Visited   bool
	Neighbors []int
}

func dfs(color, idx int, nodes []MNode, nodesMap map[int]bool) bool {
	nodes[idx].Color = color
	nodes[idx].Visited = true
	delete(nodesMap, idx)

	for _, next := range nodes[idx].Neighbors {
		if nodes[next].Visited {
			if nodes[idx].Color == nodes[next].Color {
				return false
			}
		} else {
			if !dfs(-color, next, nodes, nodesMap) {
				return false
			}
		}
	}

	return true
}

func isBipartite(graph [][]int) bool {
	nodes := make([]MNode, len(graph))
	nodesMap := make(map[int]bool)
	for i := range nodes {
		nodes[i] = MNode{0, false, graph[i]}
		nodesMap[i] = true
	}

	for len(nodesMap) > 0 {
		for idx, _ := range nodesMap {
			if !dfs(1, idx, nodes, nodesMap) {
				return false
			}

			break
		}
	}

	return true
}

