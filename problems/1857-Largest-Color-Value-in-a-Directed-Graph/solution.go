type MNode struct {
	Color     byte
	Visited   bool
	Visitable bool
	In        []int
	Out       []int
}

func traverse(idx int, nodes []MNode, colorMap map[int][]int) []int {
	if color, ok := colorMap[idx]; ok {
		return color
	}

	color := make([]int, 26)

	nodes[idx].Visited = true
	nodes[idx].Visitable = true
	if len(nodes[idx].Out) != 0 {
		for _, next := range nodes[idx].Out {
			if nodes[next].Visited {
				return nil
			} else {
				c := traverse(next, nodes, colorMap)
				if c == nil {
					return nil
				}

				for i := range color {
					if c[i] > color[i] {
						color[i] = c[i]
					}
				}
			}
		}
	}
	nodes[idx].Visited = false

	color[nodes[idx].Color]++
	colorMap[idx] = color

	return color
}

func largestPathValue(colors string, edges [][]int) int {
	nodes := make([]MNode, len(colors))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = MNode{
			Color:     colors[i] - 'a',
			Visited:   false,
			Visitable: false,
			In:        make([]int, 0),
			Out:       make([]int, 0),
		}
	}

	for _, edge := range edges {
		nodes[edge[0]].Out = append(nodes[edge[0]].Out, edge[1])
		nodes[edge[1]].In = append(nodes[edge[1]].In, edge[0])
	}

	result := -1
	colorMap := make(map[int][]int)
	for i := 0; i < len(nodes); i++ {
		if len(nodes[i].In) == 0 {
			color := traverse(i, nodes, colorMap)

			if color == nil {
				return -1
			}

			for i := range color {
				if color[i] > result {
					result = color[i]
				}
			}
		}
	}

	for i := range nodes {
		if !nodes[i].Visitable {
			return -1
		}
	}

	return result
}

