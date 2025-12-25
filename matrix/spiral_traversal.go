package matrix

// SpiralTraversal returns all elements of the matrix in spiral order.
func SpiralTraversal(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
