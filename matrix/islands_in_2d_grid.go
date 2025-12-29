package matrix

// IslandsIn2DGrid counts the number of islands (groups of connected 1s) in a 2D grid.
// Achieved Complexity: Time O(M*N), Space O(M*N) (visited array + recursion)
// Best Complexity: Time O(M*N), Space O(1) (if modification of input is allowed)
func IslandsIn2DGrid(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if visited[r][c] || matrix[r][c] == 0 {
			return
		}

		visited[r][c] = true

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	islands := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 1 && !visited[r][c] {
				islands++
				dfs(r, c)
			}
		}
	}

	return islands
}

