package arraybasics

// RotateMatrix rotates an NxN matrix 90 degrees clockwise in-place.
// Time Complexity: O(n^2) — transpose + reverse each row (touch each element a constant number of times).
// Space Complexity: O(1) — operates in-place (ignoring input/output memory).
// Best achievable: Time = O(n^2), Space = O(1) — already optimal for in-place rotation.
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
