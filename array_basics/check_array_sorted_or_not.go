package arraybasics

// isSorted checks whether the array is sorted in non-decreasing order.
// Time Complexity: O(n) — single pass.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — you must check adjacent pairs at least once.
func isSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
