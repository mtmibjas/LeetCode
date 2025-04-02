func maxVowels(s string, k int) int {
	v := []byte{'a', 'e', 'i', 'o', 'u'}

	initalCount := 0
	for i := 0; i < k; i++ {
		if slices.Contains(v, s[i]) {
			initalCount++
		}
	}
	m := initalCount
	for i := 1; i < len(s)-k+1; i++ {
		if slices.Contains(v, s[i-1]) {
			initalCount--
		}
		if slices.Contains(v, s[i+k-1]) {
			initalCount++
		}
		if m < initalCount {
			m = initalCount
		}
		if m == k {
			return m
		}
	}

	return m
}