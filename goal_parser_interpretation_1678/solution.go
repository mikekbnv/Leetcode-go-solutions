package goal_parser_interpretation_1678

import "strings"

func interpret(command string) string {
	var res strings.Builder
	command = strings.ReplaceAll(command, "()", "o")
	for _, v := range command {
		if v == '(' || v == ')' {
			continue
		}
		res.WriteRune(v)
	}
	return res.String()
}
