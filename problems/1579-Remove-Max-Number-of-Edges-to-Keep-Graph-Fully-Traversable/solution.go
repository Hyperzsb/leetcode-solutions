type UnionFind struct {
	Parent []int
}

func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{make([]int, n+1)}
	for i := 0; i <= n; i++ {
		unionFind.Parent[i] = i
	}

	return unionFind
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)

	if rootX != rootY {
		uf.Parent[rootY] = rootX
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x])
	}

	return uf.Parent[x]
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	aliceEdges, bobEdges, bothEdges := make([][]int, 0), make([][]int, 0), make([][]int, 0)
	for _, edge := range edges {
		if edge[0] == 1 {
			aliceEdges = append(aliceEdges, edge[1:])
		} else if edge[0] == 2 {
			bobEdges = append(bobEdges, edge[1:])
		} else {
			bothEdges = append(bothEdges, edge[1:])
		}
	}

	aliceUF, bobUF := NewUnionFind(n), NewUnionFind(n)

	aliceCnt, bobCnt := 0, 0
	result := 0

	for _, edge := range bothEdges {
		flag := false

		if rootX, rootY := aliceUF.Find(edge[0]), aliceUF.Find(edge[1]); rootX != rootY {
			flag = true
			aliceCnt++
			aliceUF.Union(edge[0], edge[1])
		}

		if rootX, rootY := bobUF.Find(edge[0]), bobUF.Find(edge[1]); rootX != rootY {
			flag = true
			bobCnt++
			bobUF.Union(edge[0], edge[1])
		}

		if flag {
			result++
		}
	}

	for _, edge := range aliceEdges {
		if rootX, rootY := aliceUF.Find(edge[0]), aliceUF.Find(edge[1]); rootX != rootY {
			result++
			aliceCnt++
			aliceUF.Union(edge[0], edge[1])
		}
	}

	for _, edge := range bobEdges {
		if rootX, rootY := bobUF.Find(edge[0]), bobUF.Find(edge[1]); rootX != rootY {
			result++
			bobCnt++
			bobUF.Union(edge[0], edge[1])
		}
	}

	if aliceCnt != n-1 || bobCnt != n-1 {
		return -1
	} else {
		return len(edges) - result
	}
}

