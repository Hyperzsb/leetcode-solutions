type Node struct {
	Visited   bool
	Neighbors []int
}

func traverse(idx int, nodes []Node, notVisited map[int]bool, nodeCnt *int) {
	nodes[idx].Visited = true
	*nodeCnt++
	delete(notVisited, idx)

	for _, neighbor := range nodes[idx].Neighbors {
		if !nodes[neighbor].Visited {
			traverse(neighbor, nodes, notVisited, nodeCnt)
		}
	}
}

func countPairs(n int, edges [][]int) int64 {
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{false, make([]int, 0)}
	}

	for _, edge := range edges {
		nodes[edge[0]].Neighbors = append(nodes[edge[0]].Neighbors, edge[1])
		nodes[edge[1]].Neighbors = append(nodes[edge[1]].Neighbors, edge[0])
	}

	islands := make([]int64, 0)

	notVisited := make(map[int]bool)
	for i := 0; i < n; i++ {
		notVisited[i] = true
	}

	for len(notVisited) > 0 {
		idx := 0
		for next, _ := range notVisited {
			idx = next
			break
		}

		nodeCnt := 0
		traverse(idx, nodes, notVisited, &nodeCnt)
		islands = append(islands, int64(nodeCnt))
	}

	var sum int64
	prefixSum := make([]int64, len(islands))
	for i := 0; i < len(islands); i++ {
		sum += islands[i]
		prefixSum[i] = sum
	}

	var result int64
	for i := 0; i < len(islands); i++ {
		result += (sum - prefixSum[i]) * islands[i]
	}

	return result
}

