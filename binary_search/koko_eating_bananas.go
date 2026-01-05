package binarysearch

// Time Complexity: Achieved: O(N log(max pile)), Best: O(N log(max pile))
// Space Complexity: Achieved: O(1), Best: O(1)
// KokoEatingBananas finds the minimum hourly eating speed k so Koko can eat all bananas in h hours.
// piles: number of bananas in each pile
// h: total hours available
func KokoEatingBananas(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	low, high := 1, maxPile
	answer := maxPile

	for low <= high {
		mid := low + (high-low)/2

		if canEatAll(piles, mid, h) {
			answer = mid
			high = mid - 1 
		} else {
			low = mid + 1 
		}
	}

	return answer
}

func canEatAll(piles []int, k int, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + k - 1) / k
		if hours > h {
			return false
		}
	}
	return true
}