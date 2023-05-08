func diagonalSum(mat [][]int) int {
	result := 0
	for i := 0; i < len(mat); i++ {
		result += mat[i][i] + mat[i][len(mat[0])-1-i]
	}

	if len(mat)%2 == 1 {
		result -= mat[len(mat)/2][len(mat)/2]
	}

	return result
}

