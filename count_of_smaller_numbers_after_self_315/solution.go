package count_of_smaller_numbers_after_self_315

import "sort"

func countSmaller(nums []int) []int {
	helper := make([]int, len(nums))
	copy(helper, nums)
	sort.Ints(helper)
	res := []int{}
	for _, c := range nums {
		k, i := 0, 0
		for _, v := range helper {
			if v == c {
				break
			}
			i++
			k++
		}
		helper = append(helper[:i], helper[i+1:]...)
		res = append(res, k)
	}
	return res
}
