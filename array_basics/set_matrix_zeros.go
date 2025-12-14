package arraybasics

// setMatrixZero sets entire row and column to zero if an element is zero.
// Time Complexity: O(m * n * (m + n)) in the worst case — for each zero the code updates its full row and column;
//
//	practical performance is better due to `present` guarding repeated updates.
//
// Space Complexity: O(m * n) — copies matrix and maintains a boolean `present` matrix.
func setMatrixZero(arr [][]int) [][]int {
	// Best achievable: Time = O(n*m), Space = O(1) — using first row/col as markers reduces extra space.
	var res [][]int
	present := make([][]bool, len(arr))
	for i := range arr {
		present[i] = make([]bool, len(arr[i]))
	}
	for i := range arr {
		var temp []int
		temp = append(temp, arr[i]...)
		res = append(res, temp)
	}
	for i := range arr {
		for j := range arr[i] {
			if !present[i][j] && res[i][j] == 0 {
				updateMatrix(i, j, res, present)
			}
		}
	}
	return res
}

func updateMatrix(x, y int, res [][]int, present [][]bool) {
	for i := range res {
		present[i][y] = true
		res[i][y] = 0
	}
	for i := 0; i < len(res[0]); i++ {
		present[x][i] = true
		res[x][i] = 0
	}
}
