package arraybasics

import (
	"math"
)

// second_largest_element returns the second largest distinct element in arr.
// Time Complexity: O(n) — single pass to track max and second max.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — one-pass, constant memory.
func second_largest_element(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	max := math.MinInt
	secondMax := math.MinInt

	for _, v := range arr {
		if v > max {
			secondMax = max
			max = v
		} else if v > secondMax && v != max {
			secondMax = v
		}
	}

	if secondMax == math.MinInt {
		return 0
	}

	return secondMax
}
