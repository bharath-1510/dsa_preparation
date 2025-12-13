package arraybasics

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
