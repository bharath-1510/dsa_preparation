package binarysearch

// FindPeakElement finds a peak element in an array and returns its index.
// A peak element is an element that is strictly greater than its neighbors.
// nums: array of integers
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2

        if (mid == 0 || nums[mid] >= nums[mid-1]) && (mid == len(nums)-1 || nums[mid] >= nums[mid+1]) {
            return mid
        }

        if mid > 0 && nums[mid-1] > nums[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}
