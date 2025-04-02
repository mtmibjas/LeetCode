func maxVowels(s string, k int) int {
	v := map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}}

	initalCount := 0
	for i := 0; i < k; i++ {
		if _, ok := v[s[i]]; ok {
			initalCount++
		}
	}
	m := initalCount
	for i := 1; i < len(s)-k+1; i++ {
		if  _, ok := v[s[i-1]]; ok{
			initalCount--
		}
		if  _, ok := v[s[i+k-1]]; ok {
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