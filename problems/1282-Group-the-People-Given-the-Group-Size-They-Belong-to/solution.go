func groupThePeople(groupSizes []int) [][]int {
    groups := make([][]int, len(groupSizes)+1)
    for i := range groups {
        groups[i] = make([]int, 0)
    }

    for i := range groupSizes {
        groups[groupSizes[i]] = append(groups[groupSizes[i]], i)
    }

    result := make([][]int, 0)
    for i := 1; i <= len(groupSizes); i++ {
        for j := 0; j < len(groups[i])/i; j++ {
            result = append(result, make([]int, 0))
            for k := i*j; k < i*(j+1); k++ {
                result[len(result)-1] = append(result[len(result)-1], groups[i][k])
            }
        }
    }

    return result
}

