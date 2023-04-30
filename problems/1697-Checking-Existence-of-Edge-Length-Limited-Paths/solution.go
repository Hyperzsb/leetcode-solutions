type UnionFind struct {
	Parent []int
}

func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{make([]int, n)}
	for i := 0; i < n; i++ {
		unionFind.Parent[i] = i
	}

	return unionFind
}

func (u *UnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}

	return u.Parent[x]
}

func (u *UnionFind) Union(x, y int) {
	rootX, rootY := u.Find(x), u.Find(y)

	if rootX != rootY {
		u.Parent[rootY] = rootX
	}
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(a, b int) bool {
		return edgeList[a][2] < edgeList[b][2]
	})

	idxedQueries := make([][]int, len(queries))
	for i := 0; i < len(queries); i++ {
		idxedQueries[i] = make([]int, 4)
		idxedQueries[i][0] = i
		idxedQueries[i][1] = queries[i][0]
		idxedQueries[i][2] = queries[i][1]
		idxedQueries[i][3] = queries[i][2]
	}

	sort.Slice(idxedQueries, func(a, b int) bool {
		return idxedQueries[a][3] < idxedQueries[b][3]
	})

	unionFind := NewUnionFind(n)

	result := make([]bool, len(queries))

	queryIdx, edgeIdx := 0, 0
	for queryIdx < len(idxedQueries) {
		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < idxedQueries[queryIdx][3] {
			unionFind.Union(edgeList[edgeIdx][0], edgeList[edgeIdx][1])

			edgeIdx++
		}

		if unionFind.Find(idxedQueries[queryIdx][1]) == unionFind.Find(idxedQueries[queryIdx][2]) {
			result[idxedQueries[queryIdx][0]] = true
		}

		queryIdx++
	}

	return result
}

