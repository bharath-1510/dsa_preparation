package arraymedium

// MajorityElementI returns the element that appears more than n/2 times.
//
//	func MajorityElementI(arr []int) int {
//		freq := make(map[int]int, 0)
//		for _, val := range arr {
//			freq[val]++
//		}
//		target := len(arr)/2 + 1
//		found := -1
//		for k, v := range freq {
//			if v >= target {
//				found = k
//				break
//			}
//		}
//		return found
//	}
//
// MajorityElementI returns the element that appears more than n/2 times.
// Achieved: Time Complexity: O(n), Space Complexity: O(1) — Boyer-Moore majority vote algorithm.
// Best achievable: Time Complexity: O(n), Space Complexity: O(1) — already optimal.
func MajorityElementI(arr []int) int {
	count := 0
	candidate := 0

	for _, v := range arr {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
