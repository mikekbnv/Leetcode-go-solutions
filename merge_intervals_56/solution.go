package merge_intervals_56

import "sort"

type darray [][]int

func (s darray) Len() int {
	return len(s)
}
func (s darray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s darray) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) (res [][]int) {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Sort(darray(intervals))
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] {
			tmp[1] = Max(tmp[1], intervals[i][1])
		} else {
			res = append(res, tmp)
			tmp = intervals[i]
		}
	}
	res = append(res, tmp)
	return
}
