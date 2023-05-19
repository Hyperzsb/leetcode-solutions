type MNode struct {
	Val     int
	Visited bool
	In      []int
	Out     []int
}

func traverse(idx int, nodes []MNode, nodesMap map[int]bool) {
	nodes[idx].Visited = true
	delete(nodesMap, idx)

	for _, next := range nodes[idx].Out {
		if !nodes[next].Visited {
			traverse(next, nodes, nodesMap)
		}
	}
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	nodes, nodesMap := make([]MNode, n), make(map[int]bool)
	for i := range nodes {
		nodes[i] = MNode{i, false, make([]int, 0), make([]int, 0)}
		nodesMap[i] = true
	}

	for _, edge := range edges {
		nodes[edge[0]].Out = append(nodes[edge[0]].Out, edge[1])
		nodes[edge[1]].In = append(nodes[edge[1]].In, edge[0])
	}

	result := make([]int, 0)

	for len(nodesMap) > 0 {
		for i, _ := range nodesMap {
			if len(nodes[i].In) == 0 {
				result = append(result, i)
				traverse(i, nodes, nodesMap)
				break
			}
		}
	}

	return result
}

