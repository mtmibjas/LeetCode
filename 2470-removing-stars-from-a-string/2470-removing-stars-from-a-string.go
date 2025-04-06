func removeStars(s string) string {
	stack := make([]byte, len(s))
    index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			index--
			continue
		}
		stack[index] =  s[i]
        index++
	}
	return string(stack[:index])
}
