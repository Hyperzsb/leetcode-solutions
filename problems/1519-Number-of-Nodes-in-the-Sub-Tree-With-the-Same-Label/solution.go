type MyTreeNode struct {
	Val      byte
	Visited  bool
	Children []int
}

func traverse(root int, nodes *[]MyTreeNode, labelMap *map[byte][]*int, result *[]int) {
	(*nodes)[root].Visited = true

	labelCount := 0
	(*labelMap)[(*nodes)[root].Val] = append((*labelMap)[(*nodes)[root].Val], &labelCount)
	for _, count := range (*labelMap)[(*nodes)[root].Val] {
		*count += 1
	}

	for _, child := range (*nodes)[root].Children {
		if !(*nodes)[child].Visited {
			traverse(child, nodes, labelMap, result)
		}
	}

	(*result)[root] = labelCount
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	nodes := make([]MyTreeNode, n)
	for i := 0; i < n; i++ {
		nodes[i].Val = labels[i]
	}

	for _, edge := range edges {
		nodes[edge[0]].Children = append(nodes[edge[0]].Children, edge[1])
		nodes[edge[1]].Children = append(nodes[edge[1]].Children, edge[0])
	}

	labelMap := make(map[byte][]*int)
	result := make([]int, n)
	traverse(0, &nodes, &labelMap, &result)

	return result
}

