package determine_if_string_halves_are_alike_1704

import "strings"

func halvesAreAlike(s string) bool {
	i := 0
	j := len(s) / 2
	s = strings.ToLower(s)
	c1, c2 := 0, 0
	for j < len(s) {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			c2++
		}
		if s[j] == 'a' || s[j] == 'e' || s[j] == 'i' || s[j] == 'o' || s[j] == 'u' {
			c1++
		}
		j++
		i++
	}
	if c1 == c2 {
		return true
	}
	return false
}
