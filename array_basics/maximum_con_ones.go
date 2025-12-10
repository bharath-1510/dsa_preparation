package arraybasics

func maximumConsecutiveOnes(arr []int) int {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		j := i
		for len(arr) > i && arr[i] == 1 {
			i++
		}
		if count < (i - j) {
			count = i - j
		}
	}
	return count
}
