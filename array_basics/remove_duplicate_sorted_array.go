package arraybasics

// removeDuplicatesInArray removes duplicates from a sorted array and returns a new slice.
// Time Complexity: O(n) — single pass (skipping duplicates via inner loop).
// Space Complexity: O(n) — builds and returns a new slice containing unique elements.
// Best achievable: Time = O(n), Space = O(1) — in-place two-pointer technique can remove duplicates without extra slice.
func removeDuplicatesInArray(arr []int) []int {
	var res []int
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		res = append(res, arr[i])
		for i < len(arr)-1 && arr[i] == arr[i+1] {
			i++
		}
	}
	if arr[len(arr)-1] != res[len(res)-1] {
		res = append(res, arr[len(arr)-1])
	}
	return res
}
