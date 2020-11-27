package largest_number_179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) (rez string) {
	if len(nums) == 0 {
		return ""
	}
	sort.Slice(nums[:], func(i, j int) bool {
		return (strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])) > (strconv.Itoa(nums[j]) + strconv.Itoa(nums[i]))
	})
	if nums[0] == 0 {
		return "0"
	}
	for _, i := range nums {
		rez += strconv.Itoa(i)
	}
	return
}
