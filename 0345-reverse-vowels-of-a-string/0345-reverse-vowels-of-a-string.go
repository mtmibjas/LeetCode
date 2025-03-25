func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	m := map[string]struct{}{
		"a": struct{}{},
		"e": struct{}{},
		"i": struct{}{},
		"o": struct{}{},
		"u": struct{}{},
	}
	lm := true
	rm := true
	b := []byte(s)
	for left < right {
		if _, ok := m[strings.ToLower(string(b[left]))]; ok {
		   lm = false
		}
		if _, ok := m[strings.ToLower(string(b[right]))]; ok {
			rm = false
		}
		if !lm && !rm {
			b[left], b[right] = b[right], b[left]
			lm, rm = true , true
		}
		if lm {
			left++
		} 
		if rm{
			right--
		}
	}

	return string(b)
}