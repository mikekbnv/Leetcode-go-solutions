package majotity_elements_II_229

func majorityElement(nums []int) (rez []int) {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	l := len(nums) / 3
	for k, v := range mp {
		if v > l {
			rez = append(rez, k)
		}
	}
	return
}
