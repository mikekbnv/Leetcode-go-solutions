package palindrome_partitioning_131

import "strings"

func partition(s string) [][]string {
	res := [][]string{}

	dfs([]string{}, s, &res)

	return res
}

func dfs(help []string, s string, res *[][]string) {
	l := len(strings.Join(help, ""))
	if l == len(s) {
		tmp := make([]string, len(help))
		copy(tmp, help)
		*res = append(*res, tmp)
		return
	}
	for i := l; i < len(s); i++ {
		if isPalindrome(s[l : i+1]) {
			dfs(append(help, s[l:i+1]), s, res)
		}

	}
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
