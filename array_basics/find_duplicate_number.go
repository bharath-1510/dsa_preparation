package arraybasics

// FindDuplicateNumber returns the first duplicate value found in the array.
// Time Complexity: O(n) — single pass using a hash map to detect duplicates.
// Space Complexity: O(n) — additional map to track seen values.
// Best achievable: Time = O(n), Space = O(1) — if value range allows in-place marking, else O(n) space is necessary.
func FindDuplicateNumber(arr []int) int {
	check := make(map[int]bool)
	for _, val := range arr {
		if check[val] {
			return val
		} else {
			check[val] = true
		}
	}
	return 0
}
