package asteroid_collision_735

func asteroidCollision(asteroids []int) []int {
	var ans []int
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 || len(ans) == 0 || ans[len(ans)-1] < 0 {
			ans = append(ans, asteroids[i])
		} else if ans[len(ans)-1] <= -asteroids[i] {
			if ans[len(ans)-1] < -asteroids[i] {
				i--
			}
			ans = ans[:len(ans)-1]
		}
	}
	return ans
}
