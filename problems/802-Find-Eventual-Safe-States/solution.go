type MNode struct {
    Visited bool
    Safe int
    Neighbors []int
}

func traverse(idx int, nodes []MNode) bool {
    if nodes[idx].Safe == 1 {
        return true
    }

    if nodes[idx].Safe == -1 {
        return false
    }

    nodes[idx].Visited = true

    flag := 1
    for _, neighbor := range nodes[idx].Neighbors {
        if nodes[neighbor].Visited {
            flag = -1
            break
        }

        safe := traverse(neighbor, nodes)

        if !safe {
            flag = -1
            break
        }
    }

    nodes[idx].Safe = flag

    nodes[idx].Visited = false

    if nodes[idx].Safe == 1 {
        return true
    } else {
        return false
    }
}

func eventualSafeNodes(graph [][]int) []int {
    nodes := make([]MNode, len(graph))
    for i := range nodes {
        nodes[i] = MNode{
            Neighbors: make([]int, len(graph[i])),
        }

        for j := range graph[i] {
            nodes[i].Neighbors[j] = graph[i][j]
        }
    }

    result := make([]int, 0)
    for i := range nodes {
        if safe := traverse(i, nodes); safe {
            result = append(result, i)
        }
    }

    return result
}

