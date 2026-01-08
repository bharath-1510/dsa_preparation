
package strings_fundamentals

// ZFunction computes the Z-array for string s.
func ZFunction(s string) []int {
	r := []rune(s)
	n := len(r)
	z := make([]int, n)

	l, rBound := 0, 0

	for i := 1; i < n; i++ {
		if i <= rBound {
			z[i] = min(rBound-i+1, z[i-l])
		}

		for i+z[i] < n && r[z[i]] == r[i+z[i]] {
			z[i]++
		}

		if i+z[i]-1 > rBound {
			l = i
			rBound = i + z[i] - 1
		}
	}

	return z
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

