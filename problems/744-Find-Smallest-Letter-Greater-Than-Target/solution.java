type MNode struct {
    Visited bool
    Neighbors []int
    Cnt int
}

func isInteractional(a, b []int) bool {
    distSquare := (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
    return a[2] * a[2] >= distSquare
}

func traverse(idx int, nodes []MNode, cnt *int) {
    *cnt++

    nodes[idx].Visited = true

    for _, neighbor := range nodes[idx].Neighbors {
        if !nodes[neighbor].Visited {
            traverse(neighbor, nodes, cnt)
        }
    }
}

func maximumDetonation(bombs [][]int) int {
    nodes := make([]MNode, len(bombs))
    nodeMap := make(map[int]bool)
    for i := range bombs {
        nodes[i] = MNode{
            Neighbors: make([]int, 0),
            Cnt: -1,
        }
        nodeMap[i] = true

        for j := 0; j < i; j++ {
            if isInteractional(bombs[i], bombs[j]) {
                nodes[i].Neighbors = append(nodes[i].Neighbors, j)
            }
            if isInteractional(bombs[j], bombs[i]) {
                nodes[j].Neighbors = append(nodes[j].Neighbors, i)
            }
        }
    }

    result := 0
    for i := range nodes {
        cnt := 0
        traverse(i, nodes, &cnt)

        if cnt > result {
            result = cnt
        }

        for j := range nodes {
            nodes[j].Visited = false
        }
    }

    return result
}

