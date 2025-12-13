package arraybasics

func FindRepeatingAndMissing(arr []int) (int, int) {
	dup := 0
	n := len(arr)
	sum := (n * (n + 1)) / 2
	inSum := 0
	check := make(map[int]bool)
	for _, val := range arr {
		if check[val] {
			dup = val
			continue
		}
		check[val] = true
		inSum += val
	}
	return dup, sum - inSum
}
