package arraybasics

func kadane(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := arr[0]
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = max(sum + arr[i], arr[i])
		if sum > res {
			res = sum
		}
	}
	return res
}
