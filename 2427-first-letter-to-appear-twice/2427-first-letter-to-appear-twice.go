func repeatedCharacter(s string) byte {
    arr := [26]int{}
    var r rune
    for _, c := range s {
        arr[c-'a']++
        if arr[c-'a'] == 2 {
            r = c
            break
        }
    }
   

    return byte(r)
}