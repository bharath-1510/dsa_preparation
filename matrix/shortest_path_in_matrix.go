package matrix

import (
	"container/list"
	"fmt"
)

// ShortestPathInMatrix tries ALL possible paths (brute force).
// 0 = open cell, 1 = blocked
// Moves allowed in 8 directions
// func ShortestPathInMatrix(matrix [][]int) int {
// 	n := len(matrix)
// 	if n == 0 || matrix[0][0] == 1 || matrix[n-1][n-1] == 1 {
// 		return -1
// 	}
// 	if n == 1 && matrix[0][0] == 0 {
// 		return 1
// 	}

// 	dirs := [][]int{
// 		{-1, -1}, {-1, 0}, {-1, 1},
// 		{0, -1}, {0, 1},
// 		{1, -1}, {1, 0}, {1, 1},
// 	}

// 	visited := make([][]bool, n)
// 	for i := range visited {
// 		visited[i] = make([]bool, n)
// 	}

// 	minPath := math.MaxInt64

// 	var dfs func(r, c, length int)
// 	dfs = func(r, c, length int) {
// 		if r < 0 || r >= n || c < 0 || c >= n {
// 			return
// 		}
// 		if matrix[r][c] == 1 || visited[r][c] {
// 			return
// 		}
// 		if r == n-1 && c == n-1 {
// 			if length < minPath {
// 				minPath = length + 1
// 			}
// 			return
// 		}

// 		visited[r][c] = true
// 		for _, d := range dirs {
// 			dfs(r+d[0], c+d[1], length+1)
// 		}
// 		visited[r][c] = false
// 	}

// 	dfs(0, 0, 1)

//		if minPath == math.MaxInt64 {
//			return -1
//		}
//		return minPath
//	}
// Achieved Complexity: Time O(N^2), Space O(N^2)
// Best Complexity: Time O(N^2), Space O(N^2)
func ShortestPathInMatrix(matrix [][]int) int {
	n := len(matrix)
	if n == 0 || matrix[0][0] == 1 || matrix[n-1][n-1] == 1 {
		return -1
	}

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := list.New()
	queue.PushBack([3]int{0, 0, 1})
	visited[0][0] = true

	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		cell := current.Value.([3]int)
		r, c, dist := cell[0], cell[1], cell[2]

		fmt.Printf("At cell (%d, %d) with distance %d\n", r, c, dist)

		if r == n-1 && c == n-1 {
			return dist
		}

		for _, d := range dirs {
			newR, newC := r+d[0], c+d[1]
			if newR >= 0 && newR < n && newC >= 0 && newC < n && !visited[newR][newC] && matrix[newR][newC] == 0 {
				visited[newR][newC] = true
				queue.PushBack([3]int{newR, newC, dist + 1})
			}
		}
	}

	return -1
}
