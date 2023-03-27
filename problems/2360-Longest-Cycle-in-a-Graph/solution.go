type Node struct {
	Visited int
	Step    int
	Next    int
}

func traverse(idx int, nodes []Node, step, tCnt int, notVisited map[int]bool) int {
	if nodes[idx].Visited == tCnt {
		return step - nodes[idx].Step
	}

	if nodes[idx].Visited > 0 && nodes[idx].Visited < tCnt {
		return -1
	}

	nodes[idx].Visited = tCnt
	nodes[idx].Step = step
	delete(notVisited, idx)

	if nodes[idx].Next == -1 {
		return -1
	}

	return traverse(nodes[idx].Next, nodes, step+1, tCnt, notVisited)
}

func longestCycle(edges []int) int {
	nodes := make([]Node, len(edges))
	notVisited := make(map[int]bool)
	for i := 0; i < len(edges); i++ {
		nodes[i] = Node{0, 0, edges[i]}
		notVisited[i] = true
	}

	tCnt, result := 1, -1
	for len(notVisited) > 0 {
		idx := 0
		for i, _ := range notVisited {
			idx = i
			break
		}

		length := traverse(idx, nodes, 1, tCnt, notVisited)
		if length > result {
			result = length
		}

		tCnt++
	}

	return result
}

