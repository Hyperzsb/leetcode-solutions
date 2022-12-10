func getRow(rowIndex int) []int {
	rowIndex++
	result, temp := make([]int, rowIndex), make([]int, rowIndex)

	for i := 0; i < rowIndex; i++ {
		temp[0] = 1
		temp[i] = 1

		if i == 1 {
			copy(result, temp)
		} else {
			for j := 1; j < i; j++ {
				temp[j] = result[j-1] + result[j]
			}
			copy(result, temp)
		}
	}

	return result
}

