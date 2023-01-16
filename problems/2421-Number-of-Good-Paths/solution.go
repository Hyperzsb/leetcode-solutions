func find(parent *[]int, idx int) int {
	if idx == (*parent)[idx] {
		return idx
	} else {
		(*parent)[idx] = find(parent, (*parent)[idx])
		return (*parent)[idx]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	parent := make([]int, len(vals))
	maxVal, count := make(map[int]int), make(map[int]int)
	for i := 0; i < len(vals); i++ {
		parent[i] = i
		maxVal[i] = vals[i]
		count[i] = 1
	}

	sort.Slice(edges, func(a, b int) bool {
		return max(vals[edges[a][0]], vals[edges[a][1]]) < max(vals[edges[b][0]], vals[edges[b][1]])
	})

	result := 0
	for _, edge := range edges {
		a := find(&parent, edge[0])
		b := find(&parent, edge[1])

		if maxVal[a] != maxVal[b] {
			if maxVal[a] > maxVal[b] {
				parent[b] = a
			} else {
				parent[a] = b
			}
		} else {
			parent[a] = b
			result += count[a] * count[b]
			count[b] += count[a]
		}
	}

	return result + len(vals)
}

