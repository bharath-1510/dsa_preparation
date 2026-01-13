package strings_fundamentals

// RabinKarp finds all occurrences of pattern in text using Rabin-Karp algorithm.
// Time: O((n-m+1)*m) worst, O(n + m) expected; Space: O(1) — achieved
// Best possible: O(n + m) (KMP or perfect hash)
func RabinKarp(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if m > n {
		return nil
	}

	const base = 256       // number of characters
	const mod = 1000000007 // large prime to avoid collisions

	var result []int
	if m <= 0 {
		return result
	}
	patternHash := 0
	windowHash := 0
	h := 1

	for i := 0; i < m-1; i++ {
		h = (h * base) % mod
	}

	for i := 0; i < m; i++ {
		patternHash = (base*patternHash + int(pattern[i])) % mod
		windowHash = (base*windowHash + int(text[i])) % mod
	}

	for i := 0; i <= n-m; i++ {

		if patternHash == windowHash {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				result = append(result, i)
			}
		}

		if i < n-m {
			windowHash = (base*(windowHash-int(text[i])*h) + int(text[i+m])) % mod
			if windowHash < 0 {
				windowHash += mod
			}
		}
	}

	return result
}
