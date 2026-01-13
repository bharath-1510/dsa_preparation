package strings_fundamentals

// LongestCommonPrefix finds the longest common prefix among strs.
// Time: O(S), Space: O(1) — achieved and optimal
// S = sum of all characters in all strings
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, s := range strs {
		if len(s) < len(shortest) {
			shortest = s
		}
	}

	for i := 0; i < len(shortest); i++ {
		ch := shortest[i]
		for _, s := range strs {
			if s[i] != ch {
				return shortest[:i]
			}
		}
	}

	return shortest
}
