func restoreString(s string, indices []int) string {
	m := map[int]rune{}
	for k, v := range s {
		m[indices[k]] = v
	}
	
	arr := make([]rune, len(indices))
	for i := 0; i < len(indices); i++ {
		arr[i] = m[i]
	}
	return string(arr)
}