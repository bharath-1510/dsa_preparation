package arraymedium

func CountSpecialTriplets(arr []int) int {
	n := len(arr)
	ans := 0

	for i := 0; i < n; i++ {
		x := 0
		for k := i; k < n; k++ {
			x ^= arr[k]
			if x == 0 && k > i {
				ans += k - i
			}
		}
	}

	return ans
}
