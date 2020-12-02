package plus_one_66

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i != -1; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits[:]...)
	}
	return digits
}
