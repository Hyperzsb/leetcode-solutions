func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	isVisited1, isVisited2 := make([]int, len(edges)), make([]int, len(edges))
	result1, result2 := -1, -1
	length1, length2 := 0, 0

	step1, step2 := 0, 0
	for i := node1; ; i = edges[i] {
		if isVisited1[i] != 0 {
			break
		}

		step1++
		isVisited1[i] = step1

		if edges[i] == -1 {
			break
		}
	}

	for i := node2; ; i = edges[i] {
		if isVisited2[i] != 0 {
			break
		}

		step2++
		isVisited2[i] = step2

		if isVisited1[i] != 0 {
			result2 = i
			length2 = max(isVisited1[i], isVisited2[i])
			break
		}

		if edges[i] == -1 {
			break
		}
	}

	isVisited1, isVisited2 = make([]int, len(edges)), make([]int, len(edges))

	step1, step2 = 0, 0
	for i := node2; ; i = edges[i] {
		if isVisited2[i] != 0 {
			break
		}

		step2++
		isVisited2[i] = step2

		if edges[i] == -1 {
			break
		}
	}

	for i := node1; ; i = edges[i] {
		if isVisited1[i] != 0 {
			break
		}

		step1++
		isVisited1[i] = step1

		if isVisited2[i] != 0 {
			result1 = i
			length1 = max(isVisited2[i], isVisited1[i])
			break
		}

		if edges[i] == -1 {
			break
		}
	}

	if result1 == -1 && result2 == -1 {
		return -1
	}

	if length1 < length2 {
		return result1
	} else if length1 == length2 {
		return min(result1, result2)
	} else {
		return result2
	}
}

