func convertDateToBinary(date string) string {
    d := strings.Split(date, "-")
    arr := []string{}
    for _, w := range d {
        arr = append(arr, binary(w))
    }
    return strings.Join(arr, "-")
}


func binary(date string) string {
    n, _ := strconv.Atoi(date)
    s := ""
    
    for n > 0 {
        r := n%2
        s = strconv.Itoa(r)+s
        n /= 2
    }

    return s

}