package arraybasics

func setMatrixZero(arr [][]int) [][]int {
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
