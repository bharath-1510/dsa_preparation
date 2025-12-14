package arraybasics

// arrayLinearSearch returns the index of target in arr or -1 if not found.
// Time Complexity: O(n) — linear scan.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — linear search is optimal unless data is indexed.
func arrayLinearSearch(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
