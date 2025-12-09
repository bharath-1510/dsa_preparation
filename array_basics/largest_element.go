package arraybasics

func largestElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _,val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}