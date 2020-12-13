package count_the_number_of_consistent_strings_1684

func countConsistentStrings(allowed string, words []string) int {
	help := make([]bool, 26)
	for _, c := range allowed {
		help[c-'a'] = true
	}
	res := 0
	for _, w := range words {
		y := true
		for _, c := range w {
			if help[c-'a'] == false {
				y = false
				break
			}
		}
		if y {
			res++
		}
	}
	return res
}
