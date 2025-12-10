package arraybasics

func moveZeroEnd(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	var res []int
	for i := range arr {
		if arr[i] != 0 {
			res = append(res, arr[i])
		}
	}
	n := len(res)
	for i := n; i < len(arr); i++ {
		res = append(res, 0)
	}
	return res
}
