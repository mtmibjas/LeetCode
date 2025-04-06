func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1, idx1 := changeArray(word1)
	arr2, idx2 := changeArray(word2)
    if len(idx1) != len(idx2) {
        return false
    }
	sort.Ints(arr1[:])
    sort.Ints(arr2[:])
    for k,_:= range idx1 {
        if _, ok := idx2[k]; !ok {
            return false
        }
    }

    return arr1 == arr2 
}

func changeArray(str string)([26]int, map[rune]struct{}){
    arr := [26]int{}
    m := make(map[rune]struct{})
	for _, w := range str {
		arr[w-'a']++
        m[w-'a']= struct{}{}
	}
    return arr, m
}