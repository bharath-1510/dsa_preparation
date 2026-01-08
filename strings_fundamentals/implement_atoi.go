package strings_fundamentals

// ImplementAtoi converts a string to a 32-bit signed integer.
import "math"

func ImplementAtoi(s string) int {
	runes := []rune(s)
	n := len(runes)
	i := 0
	result := 0
	for i < n && runes[i] == ' ' {
		i++
	}
	sign := 1
	if i < n {
		switch runes[i] {
		case '-':
			sign = -1
			i++
		case '+':
			i++
		}
	}

	for i < n && checkNumeric(runes[i]) {
		digit := int(runes[i] - '0')

		if result > math.MaxInt32/10 ||
			(result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}

func checkNumeric(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
