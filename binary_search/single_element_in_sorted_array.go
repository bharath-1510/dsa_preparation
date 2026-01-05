package binarysearch

// Time Complexity: Achieved: O(log N), Best: O(log N)
// Space Complexity: Achieved: O(1), Best: O(1)
// SingleElementInSortedArray finds the single element in a sorted array where every other element appears twice.
// nums: input sorted array
func SingleElementInSortedArray(nums []int) int {
	low, high := 0, len(nums)-1

    for low < high {
        mid := low + (high-low)/2

        if mid%2 == 1 {
            mid--
        }

        if nums[mid] == nums[mid+1] {
            low = mid + 2
        } else {
            high = mid
        }
    }

    return nums[low]
}
