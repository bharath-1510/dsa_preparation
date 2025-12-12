package arraybasics

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
	reverse(arr, i+1, n-1)
	return res
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
