func stringMatching(words []string) []string {
    arr := make([]string, 0)
    for i := 0; i < len(words); i++{
        for j := 0; j < len(words); j++{
            if strings.Contains(words[j], words[i]) && i != j {
               arr = append(arr, words[i])
               break
            }  
        }
    }
    return arr
}