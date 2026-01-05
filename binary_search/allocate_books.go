package binarysearch

// Time Complexity: Achieved: O(N log(sum of pages)), Best: O(N log(sum of pages))
// Space Complexity: Achieved: O(1), Best: O(1)
// AllocateBooks finds the minimum possible maximum pages assigned to a student.
// pages: array of book pages
// students: number of students
func AllocateBooks(pages []int, students int) int {
	if len(pages) < students {
		return -1
	}
	low := pages[0]
	high := 0

	for _, p := range pages {
		if p > low {
			low = p 
		}
		high += p 
	}

	result := high

	for low <= high {
		mid := low + (high-low)/2

		if isPossible(pages, students, mid) {
			result = mid
			high = mid - 1 
		} else {
			low = mid + 1 
		}
	}

	return result
}

func isPossible(pages []int, students int, maxPages int) bool {
	requiredStudents := 1
	currentSum := 0

	for _, p := range pages {
		if currentSum+p > maxPages {
			requiredStudents++
			currentSum = p

			if requiredStudents > students {
				return false
			}
		} else {
			currentSum += p
		}
	}

	return true
}
