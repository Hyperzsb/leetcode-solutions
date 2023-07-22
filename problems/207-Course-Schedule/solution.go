type MNode struct {
    Visited bool
    HasLoop int
    Neighbors []int
}

func dfs(idx int, nodes []MNode) bool {
    if nodes[idx].HasLoop == 1 {
        return true
    }

    if nodes[idx].HasLoop == -1 {
        return false
    }

    nodes[idx].Visited = true

    hasLoop := false
    for _, neighbor := range nodes[idx].Neighbors {
        if nodes[neighbor].Visited {
            hasLoop = true
            break
        }

        if dfs(neighbor, nodes) {
            hasLoop = true
        }
    }

    nodes[idx].Visited = false

    if hasLoop {
        nodes[idx].HasLoop = 1
    } else {
        nodes[idx].HasLoop = -1
    }

    return hasLoop
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    nodes := make([]MNode, numCourses)
    for i := range nodes {
        nodes[i] = MNode{
            Neighbors: make([]int, 0),
        }
    }

    for _, p := range prerequisites {
        nodes[p[1]].Neighbors = append(nodes[p[1]].Neighbors, p[0])
    }

    result := true

    for i := range nodes {
        if dfs(i, nodes) {
            result = false
            break
        }
    }

    return result
}

