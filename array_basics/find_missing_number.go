package arraybasics

func FindMissingNumber(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}
	sum := (n * (n + 1)) / 2
	for _, val := range arr {
		sum -= val
	}
	return sum
}
