package binarysearch

import "sort"

// AggressiveCows finds the largest minimum distance between cows.
// stalls: positions of stalls
// cows: number of cows

func AggressiveCows(stalls []int, cows int) int {
	sort.Ints(stalls) 
	low := 1
	high := stalls[len(stalls)-1] - stalls[0]
	result := 0

	for low <= high {
		mid := low + (high-low)/2
		if canPlaceCows(stalls, cows, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func canPlaceCows(stalls []int, cows int, dist int) bool {
	count := 1
	lastPosition := stalls[0]

	for i := 1; i < len(stalls); i++ {
		if stalls[i]-lastPosition >= dist {
			count++
			lastPosition = stalls[i]
			if count == cows {
				return true
			}
		}
	}
	return false
}
