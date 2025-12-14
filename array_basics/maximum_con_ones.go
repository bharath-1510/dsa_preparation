package arraybasics

// maximumConsecutiveOnes finds the maximum run of consecutive 1s in a binary array.
// Time Complexity: O(n) — single pass (note: current implementation increments i inside loop).
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — already optimal.
func maximumConsecutiveOnes(arr []int) int {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		j := i
		for len(arr) > i && arr[i] == 1 {
			i++
		}
		if count < (i - j) {
			count = i - j
		}
	}
	return count
}
