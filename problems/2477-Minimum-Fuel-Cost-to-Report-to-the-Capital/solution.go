type MyNode struct {
	Roads []int
}

func traverse(current, parent int, nodes *[]MyNode, seats int) (int, int64) {
	var totalNodes int = 1
	var totalFuel int64 = 0

	for _, child := range (*nodes)[current].Roads {
		if child == parent {
			continue
		}

		currentNodes, currentFuel := traverse(child, current, nodes, seats)
		totalNodes += currentNodes
		totalFuel += currentFuel
	}

	if current == 0 {
		return totalNodes, totalFuel
	}

	if totalNodes%seats != 0 {
		totalFuel += int64(totalNodes/seats + 1)
	} else {
		totalFuel += int64(totalNodes / seats)
	}

	return totalNodes, totalFuel
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	nodes := make([]MyNode, len(roads)+1)
	for i := 0; i < len(roads)+1; i++ {
		nodes[i].Roads = make([]int, 0)
	}

	for _, road := range roads {
		nodes[road[0]].Roads = append(nodes[road[0]].Roads, road[1])
		nodes[road[1]].Roads = append(nodes[road[1]].Roads, road[0])
	}

	_, totalFuel := traverse(0, -1, &nodes, seats)

	return totalFuel
}

