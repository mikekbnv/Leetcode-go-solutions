package decode_string_394

import "strconv"

func decodeString(s string) string {
	i := 0
	return helper(s, &i)
}

func helper(s string, i *int) string {
	output := ""
	for ; *i < len(s); *i = *i + 1 {
		if s[*i] >= 'a' && s[*i] <= 'z' {
			output += string(s[*i])
		} else if s[*i] >= '0' && s[*i] <= '9' {
			count := 0
			for s[*i] != '[' {
				count *= 10
				num, err := strconv.Atoi(string(s[*i]))
				if err != nil {
					panic(err)
				}
				count += num
				*i = *i + 1
			}
			*i = *i + 1
			tmp := helper(s, i)
			for j := 0; j < count; j++ {
				output += tmp
			}
		} else if s[*i] == ']' {
			return output

		}
	}
	return output
}
