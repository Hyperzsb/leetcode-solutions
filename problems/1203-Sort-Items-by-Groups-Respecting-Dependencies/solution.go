type MNode struct {
    Val int
    Visited bool
    In map[int]bool
    Out map[int]bool
}

func topologicalSort(idx int, nodes []MNode) []int {
    if len(nodes[idx].In) > 0 {
        return nil
    }

    nodes[idx].Visited = true

    result := make([]int, 1)
    result[0] = nodes[idx].Val

    for next, _ := range nodes[idx].Out {
        delete(nodes[next].In, idx)
        result = append(result, topologicalSort(next, nodes)...)
    }

    return result
}

type Group struct {
    NumMap map[int]int
    NumSlice []int
}

func internalSort(group *Group, beforeItems [][]int) bool {
    nodes := make([]MNode, len(group.NumSlice))
    for i := range nodes {
        nodes[i] = MNode{
            Val: group.NumSlice[i],
            Visited: false,
            In: make(map[int]bool),
            Out: make(map[int]bool),
        }
    }

    for i, num := range group.NumSlice {
        for _, beforeNum := range beforeItems[num] {
            if j, ok := group.NumMap[beforeNum]; ok {
                nodes[i].In[j] = true
                nodes[j].Out[i] = true 
            }
        }
    }

    newNumSlice := make([]int, 0)
    for i := range nodes {
        if !nodes[i].Visited && len(nodes[i].In) == 0 {
            newNumSlice = append(newNumSlice, topologicalSort(i, nodes)...)
        }
    }

    if len(newNumSlice) != len(group.NumSlice) {
        return false
    }

    group.NumSlice = newNumSlice
    return true
}

func externalSort(groups *[]Group, group []int, beforeItems [][]int) bool {
    nodes := make([]MNode, len(*groups))
    for i := range nodes {
        nodes[i] = MNode{
            Val: i,
            Visited: false,
            In: make(map[int]bool),
            Out: make(map[int]bool),
        }
    }

    for i := range *groups {
        for _, num := range (*groups)[i].NumSlice {
            for _, beforeNum := range beforeItems[num] {
                if j := group[beforeNum]; j != i {
                    nodes[i].In[j] = true
                    nodes[j].Out[i] = true
                }
            }
        }
    }

    newGroupsOrder := make([]int, 0)
    for i := range nodes {
        if !nodes[i].Visited && len(nodes[i].In) == 0 {
            newGroupsOrder = append(newGroupsOrder, topologicalSort(i, nodes)...)
        }
    }

    if len(newGroupsOrder) != len(*groups) {
        return false
    }

    newGroups := make([]Group, len(*groups))
    for i := range newGroups {
        newGroups[i] = (*groups)[newGroupsOrder[i]]
    }

    *groups = newGroups
    return true
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    groups := make([]Group, m)
    for i := range groups {
        groups[i] = Group{
            NumMap: make(map[int]int),
            NumSlice: make([]int, 0),
        }
    }

    for i := 0; i < n; i++ {
        if group[i] == -1 {
            groups = append(groups, Group{
                NumMap: make(map[int]int),
                NumSlice: make([]int, 1),
            })

            groups[len(groups)-1].NumMap[i] = 0
            groups[len(groups)-1].NumSlice[0] = i

            group[i] = len(groups)-1
        } else {
            groups[group[i]].NumMap[i] = len(groups[group[i]].NumSlice)
            groups[group[i]].NumSlice = append(groups[group[i]].NumSlice, i)
        }
    }

    for i := range groups {
        if !internalSort(&groups[i], beforeItems) {
            return nil
        }
    }

    if !externalSort(&groups, group, beforeItems) {
        return nil
    }

    result := make([]int, 0, n)
    for i := range groups {
        for _, num := range groups[i].NumSlice {
            result = append(result, num)
        }
    }

    return result
}

