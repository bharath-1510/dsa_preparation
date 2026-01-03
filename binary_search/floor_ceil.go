package binarysearch

// FloorCeil finds the floor (greatest <= x) and ceil (smallest >= x) in a sorted array.
// If floor or ceil doesn't exist, it returns -1 for that value.
// func FloorCeil(nums []int, x int) (floor int, ceil int) {
// 	n := len(nums)

// 	floor = -1
// 	ceil = -1

// 	low, high := 0, n-1
// 	for low <= high {
// 		mid := low + (high-low)/2
// 		if nums[mid] >= x {
// 			ceil = nums[mid]
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}

// 	low, high = 0, n-1
// 	for low <= high {
// 		mid := low + (high-low)/2
// 		if nums[mid] <= x {
// 			floor = nums[mid]
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	return floor, ceil
// }

func FloorCeil(nums []int, x int) (floor int, ceil int) {
	floor, ceil = -1, -1
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == x {
			return nums[mid], nums[mid]
		}

		if nums[mid] < x {
			floor = nums[mid]
			low = mid + 1
		} else {
			ceil = nums[mid]
			high = mid - 1
		}
	}

	return floor, ceil
}