package reformat_phone_number_1694

import "strings"

func reformatNumber(number string) string {
	var s strings.Builder
	c := 0
	check := false
	for _, v := range number {
		if v >= '0' && v <= '9' {
			check = false
			c++
			s.WriteRune(v)
			if c == 3 {
				s.WriteRune('-')
				check = true
				c = 0
			}
		}
	}
	if c == 1 {
		help := []byte(s.String())
		help[len(help)-2], help[len(help)-3] = help[len(help)-3], help[len(help)-2]
		return string(help)
	}
	if check {
		return s.String()[:len(s.String())-1]
	}
	return s.String()
}
