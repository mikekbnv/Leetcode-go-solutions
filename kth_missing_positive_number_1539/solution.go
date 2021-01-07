package kth_missing_positive_number_1539

func findKthPositive(arr []int, k int) int {
	counter := 0
	res := 0
	for i := 0; i <= len(arr)-1; i++ {
		counter = arr[i] - res - 1
		res = arr[i]
		k -= counter
		if k == 0 {
			return res - 1
		} else if k < 0 {
			return res + k - 1
		}
	}
	return arr[len(arr)-1] + k
}
