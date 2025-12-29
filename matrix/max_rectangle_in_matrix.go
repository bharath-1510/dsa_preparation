package matrix

// MaxRectangleInMatrix finds the area of the largest rectangle containing only 1s in a binary matrix.
// func MaxRectangleInMatrix(matrix [][]int) int {
//     if len(matrix) == 0 || len(matrix[0]) == 0 {
//         return 0
//     }

//     m, n := len(matrix), len(matrix[0])
//     maxArea := 0

//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == 0 {
//                 continue
//             }

//             minWidth := n
//             for k := i; k < m && matrix[k][j] == 1; k++ {
//                 width := 0
//                 for c := j; c < n && matrix[k][c] == 1; c++ {
//                     width++
//                 }

//                 if width < minWidth {
//                     minWidth = width
//                 }

//                 area := minWidth * (k - i + 1)
//                 if area > maxArea {
//                     maxArea = area
//                 }
//             }
//         }
//     }

//     return maxArea
// }

// Achieved Complexity: Time O(M*N), Space O(N)
// Best Complexity: Time O(M*N), Space O(N)
func MaxRectangleInMatrix(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    m, n := len(matrix), len(matrix[0])
    heights := make([]int, n)
    maxArea := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }

        stack := []int{}
        for k := 0; k <= n; k++ {
            curHeight := 0
            if k < n {
                curHeight = heights[k]
            }

            for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
                h := heights[stack[len(stack)-1]]
                stack = stack[:len(stack)-1]

                width := k
                if len(stack) > 0 {
                    width = k - stack[len(stack)-1] - 1
                }

                area := h * width
                if area > maxArea {
                    maxArea = area
                }
            }
            stack = append(stack, k)
        }
    }

    return maxArea
}


