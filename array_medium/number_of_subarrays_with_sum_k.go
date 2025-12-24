package arraymedium

// NumberOfSubarraysWithSumK returns the count of contiguous subarrays whose sum equals k.
// func NumberOfSubarraysWithSumK(arr []int, target int) int {
// 	ans := 0
// 	n := len(arr)
// 	for i := 0; i < n; i++ {
// 		sum := 0
// 		for j := i; j < n; j++ {
// 			sum += arr[j]
// 			if sum == target {
// 				ans++
// 			}
// 		}
// 	}
// 	return ans
// }

// NumberOfSubarraysWithSumK returns the count of contiguous subarrays whose sum equals k.
// Achieved: Time Complexity: O(n), Space Complexity: O(n) — prefix sum with hashmap.
// Best achievable: Time Complexity: O(n), Space Complexity: O(n) — this is optimal for arbitrary integers.
func NumberOfSubarraysWithSumK(arr []int, target int) int {
	count := 0
	prefixSum := 0
	freq := map[int]int{0: 1}

	for _, v := range arr {
		prefixSum += v
		if c, ok := freq[prefixSum-target]; ok {
			count += c
		}
		freq[prefixSum]++
	}

	return count
}
