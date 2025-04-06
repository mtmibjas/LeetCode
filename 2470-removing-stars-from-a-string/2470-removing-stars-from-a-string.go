func removeStars(s string) string {
    arr := []byte(s)
    index := 0
    for _, r := range arr {
        if r == '*'{
            index--
            continue
        }
        arr[index] = r
        index++
    }
    return string(arr[:index])
}