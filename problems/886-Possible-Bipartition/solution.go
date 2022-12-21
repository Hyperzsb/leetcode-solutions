func check(idx int, relations *[][]int, c int, colors *[]int) bool {
	(*colors)[idx] = c
	for i := 0; i < len((*relations)[idx]); i++ {
		if (*colors)[(*relations)[idx][i]] == 0 {
			if !check((*relations)[idx][i], relations, -c, colors) {
				return false
			}
		} else {
			if (*colors)[idx] == (*colors)[(*relations)[idx][i]] {
				return false
			}
		}
	}

	return true
}

func possibleBipartition(n int, dislikes [][]int) bool {
	relations := make([][]int, n)
	for i := 0; i < n; i++ {
		relations[i] = make([]int, 0)
	}
	for i := 0; i < len(dislikes); i++ {
		relations[dislikes[i][0]-1] = append(relations[dislikes[i][0]-1], dislikes[i][1]-1)
		relations[dislikes[i][1]-1] = append(relations[dislikes[i][1]-1], dislikes[i][0]-1)
	}

	colors := make([]int, n)

	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			if !check(i, &relations, 1, &colors) {
				return false
			}
		}
	}

	return true
}

