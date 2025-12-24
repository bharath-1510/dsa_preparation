package arraymedium

// CountSpecialTriplets counts triplets (i, j, k) with equal XORs in subarrays.
// Achieved: Time Complexity: O(n^2), Space Complexity: O(1) — nested loops with prefix XOR.
// Best achievable: Time Complexity: O(n^2), Space Complexity: O(1) — this is optimal for this problem.
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
