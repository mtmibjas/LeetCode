func maximumLengthSubstring(s string) int {
	mx := 0
	right := 0
	count := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		for m[s[i]] > 2 {
			m[s[right]]--
            right++
		}

		count++
		mx = max(mx, count-right)
	}

	return mx
}