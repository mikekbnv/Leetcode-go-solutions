package partition_labels_763

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func partitionLabels(S string) []int {
	mp := make([]int, 26)
	for i, c := range S {
		mp[c-'a'] = i
	}
	res := []int{}
	j, help := 0, 0
	for i, c := range S {
		j = max(j, mp[c-'a'])
		if j == i {
			res = append(res, i-help+1)
			help = i + 1
		}
	}
	return res
}
