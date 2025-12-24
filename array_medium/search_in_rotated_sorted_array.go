package arraymedium

// func SearchInRotatedSortedArray(arr []int, target int) int {
// 	out := -1
// 	for i, val := range arr {
// 		if val == target {
// 			return i
// 		}
// 	}
// 	return out
// }

// SearchInRotatedSortedArray returns the index of target in a rotated sorted array.
// Achieved: Time Complexity: O(log n), Space Complexity: O(1) — modified binary search.
// Best achievable: Time Complexity: O(log n), Space Complexity: O(1) — already optimal.
func SearchInRotatedSortedArray(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		if arr[low] <= arr[mid] {
			if arr[low] <= target && target < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
