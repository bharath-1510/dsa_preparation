package binarysearch

// Time Complexity: Achieved: O(N log(sum of boards)), Best: O(N log(sum of boards))
// Space Complexity: Achieved: O(1), Best: O(1)
// PaintersPartition finds the minimum time to paint all boards.
// boards: lengths of boards
// painters: number of painters
func PaintersPartition(boards []int, painters int) int {
	if len(boards) < painters {
		return -1
	}
	low, high := 0, 0
	for _, val := range boards {
		if val > low {
			low = val
		}
		high += val
	}
	result := high

	for low <= high {
		mid := low + (high-low)/2
		if isPainterAllocated(boards, painters, mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

func isPainterAllocated(boards []int, painters int, maxBoards int) bool {
	requiredPainters := 1
	currentSum := 0

	for _, board := range boards {
		if currentSum+board > maxBoards {
			requiredPainters++
			currentSum = board

			if requiredPainters > painters {
				return false
			}
		} else {
			currentSum += board
		}
	}

	return true
}
