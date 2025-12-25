package arraybasics

// rotate_array_by_right_k rotates an array right by k using reversal method.
// Time Complexity: O(n) — a few linear passes over the array.
// Space Complexity: O(n) — copies the input into a new slice.
func rotate_array_by_right_k(arr []int, k int) []int {
	// Time Complexity: O(n)
	// Space Complexity: O(n) due to copying input to res
	// Best achievable: Time = O(n), Space = O(1) — perform reversals in-place without copying.
	var res []int
	res = append(res, arr...)

	if len(res) == 0 {
		return []int{}
	}
	k = k % len(res)
	Reverse(res, 0, len(res)-1)
	Reverse(res, 0, k-1)
	Reverse(res, k, len(res)-1)
	return res
}
