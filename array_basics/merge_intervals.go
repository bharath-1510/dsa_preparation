package arraybasics

// MergeIntervals merges overlapping intervals in a list of [start,end] pairs.
// Time Complexity: O(n log n) if intervals need to be sorted first; the current
// implementation assumes input may already be sorted and runs in O(n).
// Space Complexity: O(n) — result slice stores merged intervals.
// Best achievable: Time = O(n log n) when sorting is required, Space = O(1) additional (excluding output) if merging in-place is possible.
func MergeIntervals(arr [][]int) [][]int {
	res := make([][]int, 0)
	for _, val := range arr {
		if len(res) == 0 || res[len(res)-1][1] < val[0] {
			res = append(res, val)
		} else {
			temp := res[len(res)-1]
			if temp[1] < val[1] {
				res[len(res)-1][1] = val[1]
			}
		}
	}
	return res
}
