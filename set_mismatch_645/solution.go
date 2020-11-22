package set_mismatch_645

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func findErrorNums(nums []int) []int {
	miss, dup := -1, 1
	for _, v := range nums {
		if nums[abs(v)-1] < 0 {
			dup = abs(v)
		} else {
			nums[abs(v)-1] *= -1
		}
	}

	for i, v := range nums {
		if v > 0 {
			miss = i + 1
			break
		}
	}
	return []int{dup, miss}
}
