func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, w := range patterns {

		if strings.Contains(word, w) {
			count++
			
		}

	}

	return count
}