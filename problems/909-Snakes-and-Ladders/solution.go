func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getVal(idx int, board *[][]int) int {
	n := len(*board)

	if ((idx-1)/n)%2 == 0 {
		return (*board)[n-1-(idx-1)/n][(idx-1)%n]
	} else {
		return (*board)[n-1-(idx-1)/n][n-1-(idx-1)%n]
	}
}

func setVal(idx int, board *[][]int, val int) {
	n := len(*board)

	if ((idx-1)/n)%2 == 0 {
		(*board)[n-1-(idx-1)/n][(idx-1)%n] = val
	} else {
		(*board)[n-1-(idx-1)/n][n-1-(idx-1)%n] = val
	}
}

func snakesAndLadders(board [][]int) int {
	n := len(board)

	step := make([][]int, n)
	for i := 0; i < n; i++ {
		step[i] = make([]int, n)
		for j := 0; j < n; j++ {
			step[i][j] = 10000
		}
	}

	dest := make([]int, 1, n)
	dest[0] = 1
	step[n-1][0] = 0

	for i := 0; i < len(dest); i++ {
		currentStep := getVal(dest[i], &step)

		for j := dest[i] + 1; j <= min(dest[i]+6, n*n); j++ {
			if next := getVal(j, &board); next != -1 {
				if stepVal := getVal(next, &step); stepVal > currentStep+1 {
					dest = append(dest, next)
					setVal(next, &step, currentStep+1)

					if stepVal = getVal(j, &step); stepVal > currentStep+1 {
						setVal(j, &step, currentStep+1)
					}
				}
			} else {
				if stepVal := getVal(j, &step); stepVal > currentStep+1 {
					dest = append(dest, j)
					setVal(j, &step, currentStep+1)
				}
			}
		}
	}

	if finalStep := getVal(n*n, &step); finalStep == 10000 {
		return -1
	} else {
		return finalStep
	}
}

