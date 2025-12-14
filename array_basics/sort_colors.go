package arraybasics

// sortColors sorts an array containing 0/1/2 in-place (Dutch National Flag algorithm).
// Time Complexity: O(n) — single pass with three pointers.
// Space Complexity: O(1) — sorts in-place.
// Best achievable: Time = O(n), Space = O(1) — already optimal via Dutch National Flag algorithm.
func sortColors(arr []int) []int {
	l, mid, h := 0, 0, len(arr)-1
	for mid <= h {
		switch arr[mid] {
		case 0:
			temp := arr[mid]
			arr[mid] = arr[l]
			arr[l] = temp
			mid++
			l++
		case 1:
			mid++
		case 2:
			temp := arr[mid]
			arr[mid] = arr[h]
			arr[h] = temp
			h--
		}
	}
	return arr
}
