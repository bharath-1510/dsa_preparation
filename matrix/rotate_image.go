package matrix

import (
	arraybasics "dsa_preparation/array_basics"
)

// RotateImage rotates an n x n matrix by 90 degrees clockwise in-place.
func RotateImage(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	for i := range n {
		for j := i + 1; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, val := range matrix {
		n := len(val) - 1
		arraybasics.Reverse(val, 0, n)
	}
	return matrix
}
