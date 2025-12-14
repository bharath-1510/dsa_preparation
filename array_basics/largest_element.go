package arraybasics

// largestElement finds the maximum value in the array.
// Time Complexity: O(n) — single pass over the array.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — you must examine each element once.
func largestElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}
