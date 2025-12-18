package arraymedium

// func LargestSubarrayZeroSum(arr []int) int {
	// n := len(arr)
	// maxLen := 0

	// for i := range n {
	// 	sum := 0
	// 	for j := i; j < n; j++ {
	// 		sum += arr[j]
	// 		if sum == 0 {
	// 			if j-i+1 > maxLen {
	// 				maxLen = j - i + 1
	// 			}
	// 		}
	// 	}
	// }
	// return maxLen
// }

func LargestSubarrayZeroSum(arr []int) int {
	sum := 0
	maxLen := 0
	indexMap := make(map[int]int)

	for i, val := range arr {
		sum += val

		if sum == 0 {
			maxLen = i + 1
		}

		if idx, ok := indexMap[sum]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			indexMap[sum] = i
		}
	}
	return maxLen
}

