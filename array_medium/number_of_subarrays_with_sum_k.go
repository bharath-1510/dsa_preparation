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
