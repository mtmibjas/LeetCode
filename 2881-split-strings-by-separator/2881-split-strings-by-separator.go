func splitWordsBySeparator(words []string, separator byte) []string {
	var arr []string
	for _, w := range words {
		s := 0

		for i := 0; i < len(w); i++ {
			if w[i] == separator {
				if str := string(w[s:i]); str != "" {
					arr = append(arr, str)
				}
				s = i + 1

			}else if  str := string(w[s:]); str != "" && i == len(w)-1 {
				arr = append(arr, str)
			}
		}
	}

	return arr
}