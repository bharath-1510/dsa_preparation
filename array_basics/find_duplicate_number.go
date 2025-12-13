package arraybasics

func FindDuplicateNumber(arr []int) int {
	check := make(map[int]bool)
	for _, val := range arr {
		if check[val] {
			return val
		} else {
			check[val] = true
		}
	}
	return 0
}
