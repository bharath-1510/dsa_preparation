package arraybasics

import (
	"math"
)

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
