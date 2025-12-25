package matrix

// RowWithMaxOnes finds the row with the maximum number of 1s in a binary matrix.
func RowWithMaxOnes(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	maxRow := 0
	pos := -1

	for i := 0; i < m; i++ {
		rowOnes := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				rowOnes = n - j
				break
			}
		}
		if rowOnes > maxRow {
			maxRow = rowOnes
			pos = i
		}
	}

	return pos
}
