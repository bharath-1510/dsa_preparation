package arraymedium

// func InversionCount(arr []int) int {
// 	out := 0
// 	if len(arr) < 2 {
// 		return 0
// 	}
// 	for i := range arr {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i] > arr[j] {
// 				out++
// 			}
// 		}
// 	}
// 	return out
// }

// InversionCount returns the number of inversions in the array.
// Achieved: Time Complexity: O(n log n), Space Complexity: O(n) — uses merge sort based approach.
// Best achievable: Time Complexity: O(n log n), Space Complexity: O(n) — this is optimal for inversion count.
func mergeAndCount(arr []int, left, mid, right int) int64 {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	var invCount int64 = 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			invCount += int64(mid - i + 1)
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for idx := range temp {
		arr[left+idx] = temp[idx]
	}

	return invCount
}

func mergeSortAndCount(arr []int, left, right int) int64 {
	var invCount int64 = 0
	if left < right {
		mid := (left + right) / 2
		invCount += mergeSortAndCount(arr, left, mid)
		invCount += mergeSortAndCount(arr, mid+1, right)
		invCount += mergeAndCount(arr, left, mid, right)
	}
	return invCount
}

func InversionCount(arr []int) int {
	res := mergeSortAndCount(arr, 0, len(arr)-1)
	return int(res)
}
