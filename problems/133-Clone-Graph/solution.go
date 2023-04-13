/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func traverse(node *Node, nodes map[int]*Node, visited map[int]bool) {
	if node == nil {
		return
	}

	if _, ok := nodes[node.Val]; !ok {
		nodes[node.Val] = node
	}

	visited[node.Val] = true
	for _, neighbor := range node.Neighbors {
		if !visited[neighbor.Val] {
			traverse(neighbor, nodes, visited)
		}
	}
}

func cloneGraph(node *Node) *Node {
	nodes := make(map[int]*Node)
	visited := make(map[int]bool)

	traverse(node, nodes, visited)

	newNodes := make([]*Node, len(nodes))
	for i := 0; i < len(newNodes); i++ {
		newNodes[i] = &Node{Val: i + 1, Neighbors: make([]*Node, 0)}
	}

	for i := 0; i < len(newNodes); i++ {
		for _, neighbor := range nodes[i+1].Neighbors {
			if neighbor.Val-1 < i {
				continue
			}

			newNodes[i].Neighbors =
				append(newNodes[i].Neighbors, newNodes[neighbor.Val-1])
			newNodes[neighbor.Val-1].Neighbors =
				append(newNodes[neighbor.Val-1].Neighbors, newNodes[i])
		}
	}

	result := (*Node)(nil)
	for i := 0; i < len(newNodes); i++ {
		if node.Val == i+1 {
			result = newNodes[i]
			break
		}
	}

	return result
}

