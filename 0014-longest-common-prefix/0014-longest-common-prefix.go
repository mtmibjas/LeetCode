func longestCommonPrefix(strs []string) string {
    
    prefix := strs[0]

    for i := 1; i < len(strs); i++{
        
        for len(prefix) > len(strs[i]) || prefix != strs[i][:len(prefix)] && len(prefix) > 0 {
            
           prefix =  prefix[:len(prefix)-1]
           if len(prefix) <= len(strs[i]) {
             fmt.Println(strs[i][:len(prefix)], len(prefix))
           }
        }

        if len(prefix) == 0 {
            return ""
        }
    }
    return prefix
}