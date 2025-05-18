func minimumChairs(s string) int {
	m := 0
	c := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "E" {
			c++
		} else {
			c--
		}
		if m < c {
			m = c
		}

	}

	return m
}