type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() int {
	return (*q)[0]
}

type MyNode struct {
	Red  Queue
	Blue Queue
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	nodes := make([]MyNode, n)
	for i := 0; i < n; i++ {
		nodes[i].Red = make(Queue, 0)
		nodes[i].Blue = make(Queue, 0)
	}

	for _, red := range redEdges {
		nodes[red[0]].Red.Push(red[1])
	}

	for _, blue := range blueEdges {
		nodes[blue[0]].Blue.Push(blue[1])
	}

	nodeQ := make(Queue, 0)
	colorQ := make(Queue, 0)

	for len(nodes[0].Red) > 0 {
		nodeQ.Push(nodes[0].Red.Front())
		colorQ.Push(1)
		nodes[0].Red.Pop()
	}

	for len(nodes[0].Blue) > 0 {
		nodeQ.Push(nodes[0].Blue.Front())
		colorQ.Push(-1)
		nodes[0].Blue.Pop()
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	result[0] = 0

	length := 1
	for len(nodeQ) > 0 {
		size := len(nodeQ)
		for i := 0; i < size; i++ {
			node := nodeQ.Front()
			nodeQ.Pop()
			color := colorQ.Front()
			colorQ.Pop()

			if result[node] == -1 {
				result[node] = length
			}

			if color == 1 {
				for len(nodes[node].Blue) > 0 {
					nodeQ.Push(nodes[node].Blue.Front())
					colorQ.Push(-1)
					nodes[node].Blue.Pop()
				}
			} else {
				for len(nodes[node].Red) > 0 {
					nodeQ.Push(nodes[node].Red.Front())
					colorQ.Push(1)
					nodes[node].Red.Pop()
				}
			}
		}

		length++
	}

	return result
}

