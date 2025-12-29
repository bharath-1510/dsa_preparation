package matrix

// WordSearch checks if the word exists in the board (can move horizontally/vertically, no cell reused).
// Achieved Complexity: Time O(M*N * 4^L), Space O(M*N + L) (visited array + recursion stack)
// Best Complexity: Time O(M*N * 3^L), Space O(L)
func WordSearch(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    var dfs func(i, j, idx int) bool
    dfs = func(i, j, idx int) bool {
        if idx == len(word) {
            return true
        }

        if i < 0 || i >= m || j < 0 || j >= n {
            return false
        }

        if visited[i][j] || board[i][j] != word[idx] {
            return false
        }

        visited[i][j] = true

        found := dfs(i+1, j, idx+1) ||
            dfs(i-1, j, idx+1) ||
            dfs(i, j+1, idx+1) ||
            dfs(i, j-1, idx+1)

        visited[i][j] = false
        return found
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}

