func removeStars(s string) string {
    arr := []rune(s)
    index := 0
    for _, r := range s {
        if r == '*'{
            index--
            continue
        }
        arr[index] = r
        index++
    }
    return string(arr[:index])
}