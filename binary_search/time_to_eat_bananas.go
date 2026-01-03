package binarysearch

// TimeToEatBananas finds the minimum eating speed to finish all bananas in h hours.
// piles: number of bananas in each pile
// h: total hours available
func TimeToEatBananas(piles []int, h int) int {
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

		if canFinish(piles, mid, h) {
			answer = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return answer
}

func canFinish(piles []int, speed int, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + speed - 1) / speed
		if hours > h {
			return false
		}
	}
	return true
}
