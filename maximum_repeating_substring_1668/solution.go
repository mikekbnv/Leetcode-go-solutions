package maximum_repeating_substring_1668

import "strings"

func maxRepeating(sequence string, word string) int {
	var help strings.Builder
	help.WriteString(word)
	res := 0
	for strings.Contains(sequence, help.String()) {
		res++
		help.WriteString(word)
	}
	return res
}
