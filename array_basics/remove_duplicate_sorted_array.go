package arraybasics

func removeDuplicatesInArray(arr []int) []int {
	var res []int
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		res = append(res, arr[i])
		for i < len(arr)-1 && arr[i] == arr[i+1] {
			i++
		}
	}
	if arr[len(arr)-1] != res[len(res)-1] {
		res = append(res, arr[len(arr)-1])
	}
	return res
}
