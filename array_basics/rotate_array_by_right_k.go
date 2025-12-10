package arraybasics

func rotate_array_by_right_k(arr []int,k int) []int {
	var res []int
	res = append(res, arr...)

	if len(res) == 0 {
		return []int{}
	}
	k = k % len(res)
	reverse(res, 0, len(res)-1)
	reverse(res, 0, k-1)
	reverse(res, k, len(res)-1)
	return res
}