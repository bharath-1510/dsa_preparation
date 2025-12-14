package arraybasics

// FindRepeatingAndMissing returns the repeated and missing numbers in an array that contains numbers from 1..n with one repeated and one missing.
// Time Complexity: O(n) — single pass using a hash map.
// Space Complexity: O(n) — extra map to track seen values.
// Best achievable: Time = O(n), Space = O(1) — can be reduced using mathematical formulas or index-marking when constraints allow.
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
