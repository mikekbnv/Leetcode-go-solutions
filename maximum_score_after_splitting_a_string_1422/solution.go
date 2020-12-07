package maximum_score_after_splitting_a_string_1422

func count(h1, h2 string) int {
	rez := 0
	for _, v := range h1 {
		if v == '0' {
			rez++
		}
	}
	for _, v := range h2 {
		if v == '1' {
			rez++
		}
	}
	return rez
}

func max(f, s int) int {
	if f > s {
		return f
	}
	return s
}

func maxScore(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		help1 := s[:i]
		help2 := s[i:]
		ans = max(ans, count(help1, help2))
	}

	return ans
}
