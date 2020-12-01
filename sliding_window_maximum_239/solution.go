package sliding_window_maximum_239

func MaxSlice(slice []int) (int, int) {
	res, index := -10001, 0
	for i, v := range slice {
		if v > res {
			res = v
			index = i
		}
	}
	return res, index
}
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if k == 1 {
		return nums
	}
	max, index := MaxSlice(nums[:k])
	res = append(res, max)
	for i := 1; i+k-1 < len(nums); i++ {
		if index >= i {
			max = Max(max, nums[i+k-1])
			if max <= nums[i+k-1] {
				index = i + k - 1
			}
		} else {
			max = Max(max, nums[i+k-1])
			if max > nums[i+k-1] {
				max, index = MaxSlice(nums[i : i+k])
			} else {
				index = i + k - 1
			}
		}
		res = append(res, max)
	}
	return res
}
