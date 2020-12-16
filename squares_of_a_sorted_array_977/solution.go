package squares_of_a_sorted_array_977

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r, i := 0, len(nums)-1, len(nums)-1

	for l <= r {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[i] = nums[l] * nums[l]
			i--
			l++
		} else {
			res[i] = nums[r] * nums[r]
			i--
			r--
		}
	}
	return res
}
