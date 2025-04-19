func asteroidCollision(asteroids []int) []int {
	is := true
	for is {
		is, asteroids = Collied(asteroids)
	}

	return asteroids
}

func Collied(asteroids []int) (bool, []int) {

	for i := 1; i < len(asteroids); i++ {
		if asteroids[i-1] > 0 && asteroids[i] < 0 {
			// index := i
			res := []int{}
			res = append(res, asteroids[:i-1]...)
			if math.Abs(float64(asteroids[i])) > math.Abs(float64(asteroids[i-1])) {
				res = append(res, asteroids[i])
			} else if math.Abs(float64(asteroids[i])) < math.Abs(float64(asteroids[i-1])) {
				res = append(res, asteroids[i-1])
			}

			if i+1 < len(asteroids) {
				res = append(res, asteroids[i+1:]...)
			}

			return true, res
		}
	}

	return false, asteroids
}