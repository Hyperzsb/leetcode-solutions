func traverse(stone int, isVisited *[]bool, isConnected *[][]int) {
	(*isVisited)[stone] = true

	for _, next := range (*isConnected)[stone] {
		if !(*isVisited)[next] {
			traverse(next, isVisited, isConnected)
		}
	}
}

func removeStones(stones [][]int) int {
	xMap, yMap := make(map[int][]int), make(map[int][]int)
	isVisited := make([]bool, len(stones))
	isConnected := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		isConnected[i] = make([]int, 0)

		if xStones, exist := xMap[stones[i][0]]; exist {
			for _, stone := range xStones {
				isConnected[i] = append(isConnected[i], stone)
				isConnected[stone] = append(isConnected[stone], i)
			}
			xMap[stones[i][0]] = append(xMap[stones[i][0]], i)
		} else {
			xMap[stones[i][0]] = make([]int, 1)
			xMap[stones[i][0]][0] = i
		}

		if yStones, exist := yMap[stones[i][1]]; exist {
			for _, stone := range yStones {
				isConnected[i] = append(isConnected[i], stone)
				isConnected[stone] = append(isConnected[stone], i)
			}
			yMap[stones[i][1]] = append(yMap[stones[i][1]], i)
		} else {
			yMap[stones[i][1]] = make([]int, 1)
			yMap[stones[i][1]][0] = i
		}
	}

	result := 0
	for i := 0; i < len(stones); i++ {
		if !isVisited[i] {
			traverse(i, &isVisited, &isConnected)
			result++
		}
	}

	return len(stones) - result
}

