package arraybasics

// Segregate01 moves all 1s to the end and 0s to the front (allocates new slice).
// Time Complexity: O(n) — single pass to place 1s.
// Space Complexity: O(n) — allocates a new slice of the same size.
// Best achievable: Time = O(n), Space = O(1) — can be done in-place using two pointers.
func Segregate01(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	res := make([]int, len(arr))
	j := len(arr) - 1
	for _, val := range arr {
		if val == 1 {
			res[j] = 1
			j--
		}
	}
	return res
}
