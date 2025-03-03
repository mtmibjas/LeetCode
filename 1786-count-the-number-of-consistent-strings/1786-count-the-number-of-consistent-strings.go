func countConsistentStrings(allowed string, words []string) int {
    arr := make([]string,0)
    for i := 0; i < len(allowed); i++{
        arr = append(arr, string(allowed[i]))
    }

    count := 0
    for i := 0; i < len(words); i++{
        isAllowed := true
        for j := 0; j < len(words[i]); j++{
        
            if !slices.Contains(arr, string(words[i][j])){
                isAllowed = false
                break
            }
        }
        
        if isAllowed {
            count++
        }
    }
    return count
}