func cuts(pizza [][]bool, r, c, k int, history map[string]int) int {
	key := fmt.Sprintf("%d;%d;%d", r, c, k)
	if r, ok := history[key]; ok {
		return r
	}

	if k == 1 {
		for i := r; i < len(pizza); i++ {
			for j := c; j < len(pizza[0]); j++ {
				if pizza[i][j] {
					history[key] = 1
					return 1
				}
			}
		}

		history[key] = 0
		return 0
	}

	result := 0
	rHasApple := false
	for i := r; i < len(pizza)-1 && (len(pizza)-2-i)+(len(pizza[0])-1-c)+1 >= k-1; i++ {
		if !rHasApple {
			for j := c; j < len(pizza[0]); j++ {
				if pizza[i][j] {
					rHasApple = true
					break
				}
			}

			if !rHasApple {
				continue
			}
		}

		result = (result + cuts(pizza, i+1, c, k-1, history)) % (1000000000 + 7)
	}

	cHasApple := false
	for i := c; i < len(pizza[0])-1 && (len(pizza)-1-r)+(len(pizza[0])-2-i)+1 >= k-1; i++ {
		if !cHasApple {
			for j := r; j < len(pizza); j++ {
				if pizza[j][i] {
					cHasApple = true
					break
				}
			}

			if !cHasApple {
				continue
			}
		}

		result = (result + cuts(pizza, r, i+1, k-1, history)) % (1000000000 + 7)
	}

	history[key] = result
	return result
}

func ways(pizza []string, k int) int {
	piz := make([][]bool, len(pizza))
	for i := 0; i < len(pizza); i++ {
		piz[i] = make([]bool, len(pizza[i]))
		for j := 0; j < len(pizza[0]); j++ {
			if pizza[i][j] == 'A' {
				piz[i][j] = true
			}
		}
	}

	his := make(map[string]int)

	return cuts(piz, 0, 0, k, his)
}

