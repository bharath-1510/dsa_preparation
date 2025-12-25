package arraybasics

// nextPermutation transforms arr into its next lexicographic permutation (in-place logic used).
// Time Complexity: O(n) — scan from right to find pivot, then reverse suffix.
// Space Complexity: O(n) for returned slice since code copies input into res; in-place variant would be O(1).
// Best achievable: Time = O(n), Space = O(1) — an in-place implementation is possible without extra copies.
func nextPermutation(arr []int) []int {
	var res []int
	res = append(res, arr...)
	n := len(res)
	i := n - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for arr[j] <= arr[i] {
			j--
		}
		swap(arr, i, j)
	}
	Reverse(arr, i+1, n-1)
	return res
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
