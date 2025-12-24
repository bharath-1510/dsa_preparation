package arraymedium

// func LongestSubarraySumKPositives(arr []int, k int) int {
// 	n := len(arr)
// 	maxLen := 0

// 	for i := range n {
// 		sum := 0
// 		for j := i; j < n; j++ {
// 			sum += arr[j]
// 			if sum == k {
// 				if j-i+1 > maxLen {
// 					maxLen = j - i + 1
// 				}
// 			}
// 		}
// 	}
// 	return maxLen
// }

// LongestSubarraySumKPositives returns the length of the longest subarray with sum K (positives only).
// Achieved: Time Complexity: O(n), Space Complexity: O(1) — sliding window for positive numbers.
// Best achievable: Time Complexity: O(n), Space Complexity: O(1) — this is optimal for positive integers.
func LongestSubarraySumKPositives(arr []int, k int) int {
	left, sum, maxLen := 0, 0, 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		for sum > k {
			sum -= arr[left]
			left++
		}

		if sum == k {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}
	}
	return maxLen
}
