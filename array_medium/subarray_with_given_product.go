package arraymedium

// SubarrayWithGivenProduct returns indices (start, end) of a subarray whose product is p, or -1,-1 if none.
// func SubarrayWithGivenProduct(arr []int, p int) (int, int) {
// 	n := len(arr)
// 	for i := 0; i < n; i++ {
// 		prod := 1
// 		for j := i; j < n; j++ {
// 			prod *= arr[j]
// 			if prod == p {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1
// }

// SubarrayWithGivenProduct returns indices of a subarray whose product is p.
// Achieved: Time Complexity: O(n), Space Complexity: O(1) — sliding window for positive integers.
// Best achievable: Time Complexity: O(n), Space Complexity: O(1) — already optimal for positive integers.
func SubarrayWithGivenProduct(arr []int, p int) (int, int) {
	if p == 0 {
		for i, v := range arr {
			if v == 0 {
				return i, i
			}
		}
		return -1, -1
	}

	prod := 1
	left := 0

	for right := 0; right < len(arr); right++ {
		prod *= arr[right]

		for left <= right && prod > p {
			prod /= arr[left]
			left++
		}

		if prod == p {
			return left, right
		}
	}

	return -1, -1
}
