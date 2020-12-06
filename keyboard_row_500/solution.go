package keyboard_row_500

import (
	"fmt"
	"strings"
	"unicode"
)

func findWords(words []string) []string {
	row1, row2, row3 := "qwertyuiop", "asdfghjkl", "zxcvbnm"
	var rez []string
	for _, v := range words {
		flag := true
		if strings.Contains(row1, string(unicode.ToLower(rune(v[0])))) {

			for i := 1; i < len(v); i++ {

				if !strings.Contains(row1, string(unicode.ToLower(rune(v[i])))) {
					flag = false
					break
				}
			}
		} else if strings.Contains(row2, string(unicode.ToLower(rune(v[0])))) {

			for i := 1; i < len(v); i++ {

				if !strings.Contains(row2, string(unicode.ToLower(rune(v[i])))) {
					flag = false
					break
				}
			}
		} else {
			fmt.Println(3)
			for i := 1; i < len(v); i++ {
				if !strings.Contains(row3, string(unicode.ToLower(rune(v[i])))) {
					flag = false
					break
				}
			}
		}
		if flag {
			rez = append(rez, v)
		}
	}
	return rez
}
