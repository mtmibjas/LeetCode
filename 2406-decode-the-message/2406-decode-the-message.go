func decodeMessage(key string, message string) string {
	m := make(map[rune]int)
	index := 0
	for _, e := range key {
		if _, ok := m[e]; ok  || e == ' '{
			continue
		}
		
		m[e] = index
		index++
	}

	res := ""
	for _, c := range message {
        if c == ' ' {
            res += " "
            continue
        }
		res += string(m[c] + 'a')
	}

	return res

}