package binarysearch

// KthMissingPositiveNumber finds the k-th missing positive number in a sorted array of positive integers.
// arr: sorted array of positive integers
// k: the rank of the missing number to find
func KthMissingPositiveNumber(arr []int, k int) int {
	left, right := 0, len(arr)-1
    for left <= right {
        mid := (left + right) / 2
        missingCount := arr[mid] - (mid + 1)
        if missingCount >= k {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left + k
}
