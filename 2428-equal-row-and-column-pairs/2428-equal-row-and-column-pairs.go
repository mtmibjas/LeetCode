
func equalPairs(grid [][]int) int {
	matrix := [][]int{}

	n := len(grid[0])
	for i := 0; i < n; i++ {
		m := []int{}
		for _, raw := range grid {
			m = append(m, raw[i])
		}
		matrix = append(matrix, m)
	}
	count := 0
	for _, raw := range grid {
		for _, col := range matrix {
			if slices.Equal(raw,col) {
				count++
			}
		}
	}

    return count
}