package strings_fundamentals

// ValidParentheses checks if the input string has valid matching parentheses.
// Time: O(n), Space: O(n) — achieved and optimal
func ValidParentheses(s string) bool {
	runes := []rune(s)
	stack := make([]rune, 0)
	for i := 0; i < len(runes); i += 1 {
		temp := runes[i]
		if temp == '(' || temp == '[' || temp == '{' {
			stack = append(stack, temp)
		} else {
			if len(stack) != 0 {
				top := stack[len(stack)-1]
				if (temp == ')' && top != '(') || (temp == ']' && top != '[') || (temp == '}' && top != '{') {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}
