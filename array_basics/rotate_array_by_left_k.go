package arraybasics

// rotate_array_by_left_k rotates an array left by k using three reversals.
// Time Complexity: O(n) — each reverse runs in linear time (a few passes over array).
// Space Complexity: O(n) — function creates a copy of input (res := append(res, arr...)).
// Best achievable: Time = O(n), Space = O(1) — by doing reversals in-place without copying the array.
func rotate_array_by_left_k(arr []int, k int) []int {
	var res []int
	res = append(res, arr...)

	if len(res) == 0 {
		return []int{}
	}
	k = k % len(res)
	reverse(res, 0, len(res)-1)
	reverse(res, 0, len(res)-k-1)
	reverse(res, len(res)-k, len(res)-1)
	return res
}
func reverse(res []int, start int, end int) {
	for start < end {
		temp := res[start]
		res[start] = res[end]
		res[end] = temp
		start++
		end--
	}
}
