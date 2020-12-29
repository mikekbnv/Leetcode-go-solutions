package bulls_and_cows_299

import (
	"strconv"
	"strings"
)

func getHint(secret string, guess string) string {
	same := 0
	diff := 0
	mp := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			same++
		}
		mp[secret[i]]++
	}
	for _, v := range guess {
		if i, ok := mp[byte(v)]; ok && i > 0 {
			mp[byte(v)]--
			diff++
		}
		if i, _ := mp[byte(v)]; i == 0 {
			delete(mp, byte(v))
		}
	}
	return strings.Join([]string{strconv.Itoa(same), "A", strconv.Itoa(diff - same), "B"}, "")
}
