package arraybasics

// FindMissingNumber finds the missing number in an array containing distinct numbers from 0..n with one missing.
// Time Complexity: O(n) — single pass to subtract elements from expected sum.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — you must examine each element once.
func FindMissingNumber(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}
	sum := (n * (n + 1)) / 2
	for _, val := range arr {
		sum -= val
	}
	return sum
}
