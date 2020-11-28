package bag_of_tokens_948

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, P int) int {
	score, max, l, r := 0, 0, 0, len(tokens)-1
	sort.Ints(tokens)
	flag := true
	for l <= r && flag {
		flag = false
		if P >= tokens[l] {
			score++
			P -= tokens[l]
			l++
			flag = true
		} else {
			if score > 0 {
				P += tokens[r]
				r--
				score--
				flag = true
			}
		}
		fmt.Println(score, " ", max)
		if score > max {
			max = score
		}
	}
	return max
}
