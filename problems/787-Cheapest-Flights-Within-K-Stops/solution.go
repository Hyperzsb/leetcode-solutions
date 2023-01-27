type Queue []Stop

func (q *Queue) Push(stop Stop) {
	*q = append(*q, stop)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() Stop {
	return (*q)[0]
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

type Stop struct {
	City  int
	Price int
}

type Flight struct {
	To    int
	Price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cities := make([][]Flight, n)
	for _, flight := range flights {
		cities[flight[0]] = append(cities[flight[0]], Flight{flight[1], flight[2]})
	}

	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = 10000000
	}

	stops := make(Queue, 0, n)
	stops.Push(Stop{src, 0})

	for i := 0; i <= k && !stops.Empty(); i++ {
		length := len(stops)

		for j := 0; j < length; j++ {
			stop := stops.Front()
			stops.Pop()

			for _, flight := range cities[stop.City] {
				if costs[flight.To] > stop.Price+flight.Price {
					costs[flight.To] = stop.Price + flight.Price
					stops.Push(Stop{flight.To, costs[flight.To]})
				}
			}
		}
	}

	if costs[dst] != 10000000 {
		return costs[dst]
	} else {
		return -1
	}
}

