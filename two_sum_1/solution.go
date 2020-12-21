package two_sum_1

func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	for idx, n := range nums {
		if v, ok := d[target-n]; ok {
			return []int{v, idx}
		} else {
			d[n] = idx
		}
	}
	return nil
}
