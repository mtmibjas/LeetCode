func removeStars(s string) string {
    arr := []rune{}

    for _, r := range s {
        if string(r) == "*" {
            arr = arr[:len(arr)-1]
            continue
        }
        arr = append(arr, r)
    }
    return string(arr)
}