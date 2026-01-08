package strings_fundamentals

// ImplementStrStr finds the first occurrence of needle in haystack.
func ImplementStrStr(s1, s2 string) int {
	haystack, needle := []rune(s1), []rune(s2)

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}
