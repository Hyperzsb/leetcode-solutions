type MNode struct {
	Time     int
	Children []int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func traverse(idx int, nodes []MNode) int {
	if len(nodes[idx].Children) == 0 {
		return 0
	}

	maxTime := 0
	for _, child := range nodes[idx].Children {
		maxTime = max(maxTime, traverse(child, nodes))
	}

	return nodes[idx].Time + maxTime
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	nodes := make([]MNode, n)
	for i := range nodes {
		nodes[i] = MNode{
			Time:     informTime[i],
			Children: make([]int, 0),
		}
	}

	managerIdx := 0
	for i := range manager {
		if manager[i] == -1 {
			managerIdx = i
			continue
		}

		nodes[manager[i]].Children = append(nodes[manager[i]].Children, i)
	}

	return traverse(managerIdx, nodes)
}

