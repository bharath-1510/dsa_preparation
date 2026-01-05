package binarysearch

// SplitArrayLargestSum splits an array into m non-empty subarrays to minimize the largest sum.
// nums: array of integers
// m: number of subarrays
func SplitArrayLargestSum(nums []int, m int) int {

	left, right := 0, 0
	for _, num := range nums {
		right += num
	}
	for _, p := range nums {
		if p > left {
			left = p
		}
	}
	for left <= right {
		mid := (left + right) / 2
		if canSplit(nums, m, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canSplit(nums []int, m int, maxSum int) bool {
	currentSum := 0
	subarrayCount := 1 

	for _, num := range nums {
		if currentSum+num > maxSum {
			subarrayCount++
			currentSum = num

			if subarrayCount > m {
				return false
			}
		} else {
			currentSum += num
		}
	}
	return true
}