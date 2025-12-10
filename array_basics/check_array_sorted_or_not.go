package arraybasics

func isSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	for i:=0;i < len(arr)-1;i++ {
		if arr[i] > arr[i+1]{
			return false
		}
	}
	return true
}