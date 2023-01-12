type MyTreeNode struct {
	Apple    bool
	Needed   bool
	Visited  bool
	Children []int
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	nodes := make([]MyTreeNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = MyTreeNode{
			Apple:    hasApple[i],
			Visited:  false,
			Children: make([]int, 0),
		}
	}

	for _, edge := range edges {
		nodes[edge[0]].Children = append(nodes[edge[0]].Children, edge[1])
		nodes[edge[1]].Children = append(nodes[edge[1]].Children, edge[0])
	}

	stack := make([]int, 1, n)
	stack[0] = 0
	nodes[0].Needed = true

	result := 0
	for len(stack) > 0 {
		current := &(nodes[stack[len(stack)-1]])

		if !current.Visited && current.Apple {
			length := 0
			for i := len(stack) - 1; i >= 0; i-- {
				if nodes[stack[i]].Needed {
					result += length * 2
					break
				} else {
					nodes[stack[i]].Needed = true
				}
				length++
			}
		}

		current.Visited = true

		flag := true
		for _, child := range current.Children {
			if !nodes[child].Visited {
				flag = false
				stack = append(stack, child)
				break
			}
		}

		if flag {
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

