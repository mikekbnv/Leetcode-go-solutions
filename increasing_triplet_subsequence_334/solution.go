package increasing_triplet_subsequence_334

func increasingTriplet(nums []int) bool {
	n1, n2 := 1<<63-1, 1<<63-1
	for _, c := range nums {
		if c <= n1 {
			n1 = c
		} else if c <= n2 {
			n2 = c
		} else {
			return true
		}
	}
	return false
}
