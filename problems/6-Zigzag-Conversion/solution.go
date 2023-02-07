func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		index := i % (2*numRows - 2)
		if index >= numRows {
			index = (2*numRows - 2) - index
		}

		rows[index] = append(rows[index], s[i])
	}

	result := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		result = append(result, rows[i]...)
	}

	return string(result)
}

