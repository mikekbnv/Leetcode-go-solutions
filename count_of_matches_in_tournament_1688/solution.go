package count_of_matches_in_tournament_1688

func numberOfMatches(n int) int {
	res := 0
	for n != 1 {
		if n%2 == 0 {
			res = res + n/2
			n /= 2
		} else {
			res = res + (n-1)/2
			n = (n-1)/2 + 1
		}
	}
	return res
}
