package valid_anagrama_242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp := make([]int, 26)
	for i := 0; i < len(t); i++ {
		mp[s[i]-'a']++
		mp[t[i]-'a']--

	}

	for _, k := range mp {
		if k != 0 {
			return false
		}
	}
	return true
}
