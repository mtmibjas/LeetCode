func findWordsContaining(words []string, x byte) []int {
    arr := []int{}
    for i, word := range words {
        isExist := false    
        for i,j := 0, len(word)-1; i <= j; i, j = i+1,j-1{
            if word[i] == x || word[j] == x {
                isExist = true
                break
            }
        }
        if !isExist {
            continue
        }
        arr = append(arr,i)
    }
    return arr
}