package arraybasics

func Segregate01(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	res := make([]int, len(arr))
	j := len(arr) - 1
	for _, val := range arr {
		if val == 1 {
			res[j] = 1
			j--
		}
	}
	return res
}
