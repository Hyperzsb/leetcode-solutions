type UnionFind struct {
	Parent []int
	Rank   []int
	Cnt    int
}

func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{make([]int, n), make([]int, n), n}
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

	if rootX == rootY {
		return
	}

	if u.Rank[rootX] < u.Rank[rootY] {
		rootX ^= rootY
		rootY ^= rootX
		rootX ^= rootY
	}

	u.Parent[rootY] = rootX

	if u.Rank[rootX] == u.Rank[rootY] {
		u.Rank[rootX]++
	}

	u.Cnt--
}

func isSimilar(a, b string) bool {
	if a == b {
		return true
	}

	diffCnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffCnt++

			if diffCnt > 2 {
				return false
			}
		}
	}

	return true
}

func numSimilarGroups(strs []string) int {
	unionFind := NewUnionFind(len(strs))

	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				unionFind.Union(i, j)
			}
		}
	}

	return unionFind.Cnt
}

