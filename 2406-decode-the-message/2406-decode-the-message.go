func decodeMessage(key string, message string) string {
	m := make(map[rune]rune)
	index  := rune(0)
	for _, e := range key {
		if _, ok := m[e]; ok || e == 32{
			continue
		}
		m[e] = index + 'a'
		index++
	}
    m[32] = 32
	res := make([]rune, len(message)) 
	for i, c := range message {
		res[i] = m[c] 
	}
    
	return string(res)

}