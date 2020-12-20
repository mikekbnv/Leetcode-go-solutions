package four_sum_II_454

func fourSumCount(A, B, C, D []int) int {
	mp1 := make(map[int]int)

	for _, n := range A {
		for _, c := range B {
			mp1[c+n]++
		}
	}
	res := 0
	for _, n := range C {
		for _, c := range D {
			if k, _ := mp1[-(c + n)]; k > 0 {
				res += k
			}
		}
	}
	return res

}
