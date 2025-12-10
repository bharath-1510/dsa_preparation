package arraybasics

func rotate_array_by_left_k(arr []int, k int) []int {
	var res []int
	res = append(res, arr...)

	if len(res) == 0 {
		return []int{}
	}
	k = k % len(res)
	reverse(res, 0, len(res)-1)
	reverse(res, 0, len(res)-k-1)
	reverse(res, len(res)-k, len(res)-1)
	return res
}
func reverse(res []int, start int, end int) {
	for start < end {
		temp := res[start]
		res[start] = res[end]
		res[end] = temp
		start++
		end--
	}
}
