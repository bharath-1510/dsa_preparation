package arraymedium

// func LongestSubarraySumK(arr []int, k int) int {
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

func LongestSubarraySumK(arr []int,k int) int {
	sum := 0
	maxLen := 0
	indexMap := make(map[int]int)

	for i, val := range arr {
		sum += val

		if sum == k {
			maxLen = i + 1
		}

		if idx, ok := indexMap[sum-k]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}
		if _, exists := indexMap[sum]; !exists {
			indexMap[sum] = i
		}
	}
	return maxLen
}
