package arraybasics

// MergeSortedArrays merges two sorted arrays into one sorted array.
// Time Complexity: O(n + m) — single pass merging both arrays.
// Space Complexity: O(n + m) — result slice contains all elements from both inputs.
// Best achievable: Time = O(n + m), Space = O(1) if merging into one of the inputs in-place (if capacity permits), otherwise O(n + m).
func MergeSortedArrays(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}
