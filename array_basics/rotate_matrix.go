package arraybasics

func RotateMatrix(arr [][]int) [][]int {
	m := len(arr)
	n := len(arr[0])
	for i := range n {
		for j := i + 1; j < m; j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	for _, val := range arr {
		n := len(val) - 1
		reverse(val, 0, n)
	}
	return arr
}
