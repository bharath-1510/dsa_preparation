package strings_fundamentals

// MinimumWindowSubstring finds the minimum window in s containing all characters of t.
// Time: O(n), Space: O(1) — achieved and optimal
func MinimumWindowSubstring(s, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(t)
	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		// Include current character
		if need[s[right]] > 0 {
			required--
		}
		need[s[right]]--

		for required == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			need[s[left]]++
			if need[s[left]] > 0 {
				required++
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}
