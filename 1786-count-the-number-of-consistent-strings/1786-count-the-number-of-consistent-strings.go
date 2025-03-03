func countConsistentStrings(allowed string, words []string) int {

	count := 0
loop:
	for i := 0; i < len(words); i++ {

		for j := 0; j < len(words[i]); j++ {
			if !strings.Contains(allowed, string(words[i][j])) {
				continue loop
			}
		}

		count++

	}
	return count
}