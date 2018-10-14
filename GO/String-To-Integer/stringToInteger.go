package stringToInteger

// Implement atoi which converts a string to an integer.

// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

// If no valid conversion could be performed, a zero value is returned.

// Note:

// Only the space character ' ' is considered as whitespace character.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}
	sign, body := getBody(s)
	return checkBody(body, sign)
}

func getBody(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		sign = -1
		s = s[1:]
	case '+':
		s = s[1:]
	default:
	}
	return sign, s
}

func checkBody(s string, sign int) int {
	res := 0
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]) - 48
			tmp := res * sign
			if tmp > math.MaxInt32 {
				return math.MaxInt32
			}
			if tmp < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return res * sign
}
