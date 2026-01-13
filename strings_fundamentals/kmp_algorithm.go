package strings_fundamentals

// KMPAlgorithm finds all occurrences of pattern in text using KMP.
// Time: O(n + m), Space: O(m) — achieved and optimal
func KMPAlgorithm(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return nil
	}

	lps := make([]int, m)
	for i, j := 1, 0; i < m; {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else if j > 0 {
			j = lps[j-1]
		} else {
			lps[i] = 0
			i++
		}
	}

	var result []int
	for i, j := 0, 0; i < n; {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j)
			j = lps[j-1]
		} else if i < n && text[i] != pattern[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return result
}
