func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := [26]int{}
	for _, w := range word1 {
		arr1[w-'a']++
	}
	arr2 := [26]int{}
	for _, w := range word2 {
		arr2[w-'a']++
	}

	for i, n := range arr1 {
		if n == 0 {
			continue
		}
        k := arr2[i]
        if k == 0 {
            return false
        }
		v := slices.Index(arr2[:], n)
		if v == -1 {
			return false
		}
		arr2[v] = -2

	}
	for _, n := range arr2 {
		if n == 0 || n == -2{
			continue
		}
		return false

	}
	return true
}