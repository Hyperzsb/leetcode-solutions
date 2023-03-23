type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Front() int {
	return (*q)[0]
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

type City struct {
	Visited bool
	Roads   map[int]int
}

func minScore(n int, roads [][]int) int {
	cities := make([]City, n)
	for i := 0; i < n; i++ {
		cities[i] = City{Visited: false, Roads: make(map[int]int)}
	}
	for _, road := range roads {
		cities[road[0]-1].Roads[road[1]-1] = road[2]
		cities[road[1]-1].Roads[road[0]-1] = road[2]
	}

	queue := make(Queue, 0)
	queue.Push(0)

	result := 100000
	for len(queue) > 0 {
		curr := queue.Front()
		queue.Pop()

		for next, dist := range cities[curr].Roads {
			if !cities[next].Visited {
				cities[next].Visited = true
				queue.Push(next)
			}

			if dist < result {
				result = dist
			}
		}
	}

	return result
}

