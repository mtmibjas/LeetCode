func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	sum := 0
	for _, num := range nums {
		m[num]++
		c := m[num]
		if c == 1 {
			sum += num
		} else if c == 2 {
			sum -= num
		}
	}
    return sum
}