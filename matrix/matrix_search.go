package matrix

// MatrixSearch searches for a value in a matrix with each row and column sorted in ascending order.
func MatrixSearch(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] <= target && matrix[i][n-1] >= target {
			for j := 0; j < n; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}
