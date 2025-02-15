func replaceElements(arr []int) []int {
    mx := -1
    for i := len(arr)-1 ; i >= 0; i--{
        tmx := mx
        mx = max(mx, arr[i])
        arr[i] = tmx
    }
    return arr
}