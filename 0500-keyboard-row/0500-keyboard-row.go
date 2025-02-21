func findWords(words []string) []string {
	m := make(map[string]int)
	top := "qwertyuiop"
	for i := 0; i < len(top); i++ {
		m[string(top[i])] = 1
	}
	mid := "asdfghjkl"
	for i := 0; i < len(mid); i++ {
		m[string(mid[i])] = 2
	}
	bot := "zxcvbnm"
	for i := 0; i < len(bot); i++ {
		m[string(bot[i])] = 3
	}
	arr := make([]string, 0)
	for _, word := range words {
		isTrue := true
		f, _ := m[strings.ToLower(string(word[0]))]
		for i := 1; i < len(word); i++ {
            letter := strings.ToLower(string(word[i]))
			v, _ := m[letter]
			if v != f {
				isTrue = false
				break
			}
		}

		if isTrue {
			arr = append(arr, word)
		}
	}
	return arr
}