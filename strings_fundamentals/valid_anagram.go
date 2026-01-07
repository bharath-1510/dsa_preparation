package strings_fundamentals

import "slices"

// ValidAnagram checks if two strings are anagrams.
func ValidAnagram(s, t string) bool {
	rune1, rune2 := []rune(s), []rune(t)
	slices.Sort(rune1)
	slices.Sort(rune2)
	for i := range rune1 {
		if rune1[i] != rune2[i] {
			return false
		}
	}

	return true
}
