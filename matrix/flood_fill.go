package matrix

// FloodFill fills all connected cells of the same color as the start cell with the new color.
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	origColor := image[sr][sc]
	if origColor == newColor {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
			return
		}
		if image[r][c] != origColor {
			return
		}

		image[r][c] = newColor

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(sr, sc)
	return image
}

