package smallest_range_II_910

import "sort"

func smallestRangeII(A []int, K int) int {
	n := len(A) - 1
	sort.Ints(A)
	ans := A[n] - A[0]
	for i := 0; i < n; i++ {
		ans = Min(Max(A[n]-K, A[i]+K)-Min(A[0]+K, A[i+1]-K), ans)
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
