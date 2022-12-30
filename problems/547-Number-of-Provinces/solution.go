func traverse(city int, isVisited *[]bool, isConnected *[][]int) {
	(*isVisited)[city] = true

	for i := 0; i < len(*isConnected); i++ {
		if (*isConnected)[city][i] == 1 && !(*isVisited)[i] {
			traverse(i, isVisited, isConnected)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	cityNum, provinceNum := len(isConnected), 0
	isVisited := make([]bool, cityNum)
	for i := 0; i < cityNum; i++ {
		if !isVisited[i] {
			traverse(i, &isVisited, &isConnected)
			provinceNum++
		}
	}

	return provinceNum
}

