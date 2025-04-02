func longestCommonPrefix(strs []string) string {
    com := strs[0]

    for len(com) > 0 {
        isExist := true
        for i := 1; i < len(strs); i++{
            if len(com) > len(strs[i]) || !strings.Contains(strs[i][:len(com)], com) {
                isExist = false
                break
            }
        }
        if isExist{
            return com
        }
        com  = com[:len(com)-1]
    }

    return ""
}