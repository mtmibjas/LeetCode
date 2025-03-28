func commonChars(words []string) []string {
  
    c := []string{}
    for _, w := range words[0] {
        c = append(c , string(w))
    } 
    com := []string{}
    arr := words[1:]
    for _, s := range c {
        isCom := true
        for i, str := range arr{
            v := strings.Index(str, s)
            if v == -1 {
                isCom = false
                break
            }
            if v == len(str) -1 {
                arr[i] = str[:v]
            }else{
                arr[i] = str[:v]+str[v+1:]
            } 
        }
        if isCom {
            com = append(com, s)
        }
    }

    return com
    
}