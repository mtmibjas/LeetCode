func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")

	for i := 0; i < len(arr); i++ {
		curr := arr[i][len(arr[i])-1]
		if i < len(arr)-1 {
			next := arr[i+1][0]
			if curr != next {
				return false
			}
            continue
		}
		next := arr[0][0]
		if curr != next {
			return false
		}

	}

	return true
}