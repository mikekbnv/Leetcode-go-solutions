package find_first_and_last_position_of_element_in_sorted_array_34

func searchRange(nums []int, target int) []int {
    var p = []int{-1, -1}
    for i := 0; i < len(nums); i++{
        if nums[i] == target {
            p[0] = i
            for k := i; k < len(nums); k++{
                if nums[k] != target {
                   p[1] = k - 1
                    return p
                }
            }
            break
        } 
    }
    if len(nums) == 0 {
        return p
    }
    if nums[len(nums) - 1] == target {
        p[1] = len(nums) - 1
    }
    return p
}