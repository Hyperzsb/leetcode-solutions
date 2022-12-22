type Node struct {
	Parent   int
	Children []int
	CSum     int
	PSum     int
}

func build(root int, nodes *[]Node, edges *[]map[int]bool) {
	if len((*edges)[root]) == 0 {
		return
	}
	for key, _ := range (*edges)[root] {
		(*nodes)[root].Children = append((*nodes)[root].Children, key)
		(*nodes)[key].Parent = root
		delete((*edges)[key], root)
		build(key, nodes, edges)
	}
}

func traverse(root int, nodes *[]Node) {
	if len((*nodes)[root].Children) == 0 {
		return
	}
	for i := 0; i < len((*nodes)[root].Children); i++ {
		traverse((*nodes)[root].Children[i], nodes)
		(*nodes)[root].CSum += (*nodes)[(*nodes)[root].Children[i]].CSum + 1
		(*nodes)[root].PSum += (*nodes)[(*nodes)[root].Children[i]].PSum + (*nodes)[(*nodes)[root].Children[i]].CSum + 1
	}
}

func calc(n, root int, nodes *[]Node, result *[]int) {
	(*result)[root] = (*nodes)[root].PSum
	for _, child := range (*nodes)[root].Children {
		(*nodes)[child].PSum = (*nodes)[root].PSum - ((*nodes)[child].CSum + 1) + n - ((*nodes)[child].CSum + 1)
		calc(n, child, nodes, result)
	}
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	nodeEdges := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		nodeEdges[i] = make(map[int]bool)
	}
	for i := 0; i < len(edges); i++ {
		nodeEdges[edges[i][0]][edges[i][1]] = true
		nodeEdges[edges[i][1]][edges[i][0]] = true
	}

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].Parent = -1
		nodes[i].Children = make([]int, 0)
	}

	build(0, &nodes, &nodeEdges)

	traverse(0, &nodes)

	result := make([]int, n)
	calc(n, 0, &nodes, &result)

	return result
}

