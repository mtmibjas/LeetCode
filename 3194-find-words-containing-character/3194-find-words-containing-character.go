func findWordsContaining(words []string, x byte) []int {
    arr := []int{}
    for i, word := range words {
        if !strings.Contains(word, string(x)){
            continue
        }
        arr = append(arr,i)
    }
    return arr
}