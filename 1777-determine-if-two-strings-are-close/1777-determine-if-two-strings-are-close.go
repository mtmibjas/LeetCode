func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1, idx1 := changeArray(word1)
	arr2, idx2 := changeArray(word2)
   
	sort.Ints(arr1[:])
    sort.Ints(arr2[:])
    
    return arr1 == arr2 && idx1 == idx2
}

func changeArray(str string)([26]int, [26]int){
    arr := [26]int{}
    idx := [26]int{}
	for _, w := range str {
		arr[w-'a']++
       idx[w-'a'] =1
	}
    return arr, idx
}