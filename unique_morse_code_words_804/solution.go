package unique_morse_code_words_804

import "strings"

func uniqueMorseRepresentations(words []string) int {
	morses := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	mp := make(map[string]int)

	for _, v := range words {
		tmp := strings.Builder{}
		for _, ch := range v {
			tmp.WriteString(morses[ch-'a'])
		}
		mp[tmp.String()]++
	}
	return len(mp)
}
