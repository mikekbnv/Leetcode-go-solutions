package basic_calculator_II_227

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.ReplaceAll(s, " ", "")
	prev, curr, sign, ans := 0, 0, rune('+'), 0
	for i, v := range s {
		if v >= '0' && v <= '9' {
			curr = curr*10 + int(v-'0')
		}

		if (v < '0' || v > '9') || i == len(s)-1 {
			if sign == '+' {
				ans = ans + prev
				prev = curr
			} else if sign == '-' {
				ans = ans + prev
				prev = -curr
			} else if sign == '*' {
				fmt.Println("1")
				prev = prev * curr
			} else if sign == '/' {
				prev = prev / curr
			}

			sign = rune(s[i])
			curr = 0
		}
	}
	ans = ans + prev
	return ans
}
