type MyTreeNode struct {
	Val      byte
	Children []int
}

func traverse(root int, nodes *[]MyTreeNode, result *int) int {
	if len((*nodes)[root].Children) == 0 {
		return 1
	}

	max1, max2 := 0, 0
	for _, child := range (*nodes)[root].Children {
		if (*nodes)[child].Val != (*nodes)[root].Val {
			length := traverse(child, nodes, result)
			if length > max1 {
				max2 = max1
				max1 = length
			} else if length < max2 {
				continue
			} else {
				max2 = length
			}
		} else {
			traverse(child, nodes, result)
		}
	}

	if max1+1+max2 > *result {
		*result = max1 + 1 + max2
	}

	return max1 + 1
}

func longestPath(parent []int, s string) int {
	nodes := make([]MyTreeNode, len(parent))
	for i := 0; i < len(parent); i++ {
		nodes[i].Val = s[i]
		if parent[i] != -1 {
			nodes[parent[i]].Children = append(nodes[parent[i]].Children, i)
		}
	}

	result := 1
	traverse(0, &nodes, &result)

	return result
}

