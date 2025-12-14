package arraybasics

// moveZeroEnd moves all zeros to the end while preserving order of non-zero elements.
// Time Complexity: O(n) — single pass to collect non-zeros then append zeros.
// Space Complexity: O(n) — constructs a new slice for result.
// Best achievable: Time = O(n), Space = O(1) — can be done in-place with two pointers.
func moveZeroEnd(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	var res []int
	for i := range arr {
		if arr[i] != 0 {
			res = append(res, arr[i])
		}
	}
	n := len(res)
	for i := n; i < len(arr); i++ {
		res = append(res, 0)
	}
	return res
}
