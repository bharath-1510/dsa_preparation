package matrix


// func BooleanMatrix(matrix1 [][]int) [][]int {
// 	m := len(matrix1)
// 	n := len(matrix1[0])

// 	matrix := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		matrix[i] = make([]int, n)
// 		copy(matrix[i], matrix1[i])
// 	}

// 	visited := make([][]bool, m)
// 	for i := 0; i < m; i++ {
// 		visited[i] = make([]bool, n)
// 	}

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if !visited[i][j] && matrix[i][j] == 1 {
// 				updateValue(i, j, m, n, matrix, visited)
// 			}
// 		}
// 	}
// 	return matrix
// }

// func updateValue(i, j, m, n int, matrix [][]int, visited [][]bool) {
// 	for col := 0; col < n; col++ {
// 		matrix[i][col] = 1
// 		visited[i][col] = true
// 	}
// 	for row := 0; row < m; row++ {
// 		matrix[row][j] = 1
// 		visited[row][j] = true
// 	}
// }

// Achieved Complexity: Time O(M*N), Space O(M+N)
// Best Complexity: Time O(M*N), Space O(1)
func BooleanMatrix(matrix [][]int) [][]int {
    m := len(matrix)
    n := len(matrix[0])

    rowFlag := make([]bool, m)
    colFlag := make([]bool, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                rowFlag[i] = true
                colFlag[j] = true
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rowFlag[i] || colFlag[j] {
                matrix[i][j] = 1
            }
        }
    }

    return matrix
}
