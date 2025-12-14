package arraybasics

// kadane returns the maximum subarray sum using Kadane's algorithm.
// Time Complexity: O(n) — single pass tracking current sum and max.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — already optimal.
func kadane(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := arr[0]
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = max(sum+arr[i], arr[i])
		if sum > res {
			res = sum
		}
	}
	return res
}
