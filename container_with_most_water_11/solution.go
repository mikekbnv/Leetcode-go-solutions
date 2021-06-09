package container_with_most_water

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxArea(h []int) int {
	l, r := 0, len(h)-1
	maxarr := 0
	for l < r {
		maxarr = max(maxarr, (r-l)*min(h[l], h[r]))

		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return maxarr
}
