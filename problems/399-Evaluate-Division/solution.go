type UnionFind struct {
	Parent   []int
	Children []int
	Factor   []float64
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		Parent:   make([]int, n),
		Children: make([]int, 0),
		Factor:   make([]float64, n),
	}

	for i := 0; i < n; i++ {
		uf.Parent[i] = i
		uf.Factor[i] = 1.0
	}

	return &uf
}

func (uf *UnionFind) Union(x, y int, factor float64) {
	rootX, rootY := uf.Find(x), uf.Find(y)

	if rootX == rootY {
		return
	}

	uf.Parent[rootY] = rootX
	uf.Factor[rootY] = factor * uf.Factor[x] / uf.Factor[y]
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		oldParent := uf.Parent[x]
		uf.Parent[x] = uf.Find(uf.Parent[x])
		uf.Factor[x] *= uf.Factor[oldParent]
	}

	return uf.Parent[x]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	stringMap := make(map[string]int)
	for _, equation := range equations {
		if _, ok := stringMap[equation[0]]; !ok {
			stringMap[equation[0]] = len(stringMap)
		}

		if _, ok := stringMap[equation[1]]; !ok {
			stringMap[equation[1]] = len(stringMap)
		}
	}

	uf := NewUnionFind(len(stringMap))

	for i, equation := range equations {
		uf.Union(stringMap[equation[0]], stringMap[equation[1]], values[i])
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := stringMap[query[0]]; !ok {
			result[i] = -1.0
			continue
		}

		if _, ok := stringMap[query[1]]; !ok {
			result[i] = -1.0
			continue
		}

		rootX, rootY := uf.Find(stringMap[query[0]]), uf.Find(stringMap[query[1]])
		if rootX != rootY {
			result[i] = -1.0
		} else {
			result[i] = uf.Factor[stringMap[query[1]]] / uf.Factor[stringMap[query[0]]]
		}
	}

	return result
}

