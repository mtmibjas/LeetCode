func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := changeArray(word1)
	arr2 := changeArray(word2)

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

func changeArray(str string)[26]int{
    arr := [26]int{}
	for _, w := range str {
		arr[w-'a']++
	}
    return arr
}