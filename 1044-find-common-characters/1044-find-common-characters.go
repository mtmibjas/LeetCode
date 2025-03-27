func commonChars(words []string) []string {
  
    arr := make([][]string,len(words))

    for i, word := range words {
        a := make([]string, len(word))
        for j, w := range word {
            a[j] = string(w)
        }
        arr[i] = a
    }
    if len(arr) ==1 {
        return arr[0]
    }

    m := []string{}
    for _, c := range arr[0]{
        is := true
        for _,word := range arr[1:] {
            v := slices.Index(word, string(c))
            if v == -1 {
                is = false
                break
            }
            word[v] = ""
        }
        if is {
            m = append(m, string(c))
        }
    }

    return m
    
}