func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	midRow, midCol := row/2, 0

	if target < matrix[midRow][0] {
		if midRow > 0 {
			return searchMatrix(matrix[:midRow], target)
		} else {
			return false
		}
	} else if target > matrix[midRow][col-1] {
		if midRow < row-1 {
			return searchMatrix(matrix[midRow+1:], target)
		} else {
			return false
		}
	} else {
		start, end := 0, col-1
		for start < end {
			midCol = (start + end) / 2
			if matrix[midRow][midCol] == target {
				return true
			} else if matrix[midRow][midCol] > target {
				end = midCol
			} else {
				start = midCol + 1
			}
		}

		midCol = (start + end) / 2
		if matrix[midRow][midCol] == target {
			return true
		} else {
			return false
		}
	}
}

