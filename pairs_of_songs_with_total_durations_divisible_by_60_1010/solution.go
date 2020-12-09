package pairs_of_songs_with_total_durations_divisible_by_60_1010

func numPairsDivisibleBy60(time []int) int {
    help := make([]int, 60)
    res := 0
    for _,t := range time {
        if t % 60 == 0 {
            res += help[0]
        }else {
            res += help[60 - t % 60]
        }
        help[t % 60]++
    }
    return res
}