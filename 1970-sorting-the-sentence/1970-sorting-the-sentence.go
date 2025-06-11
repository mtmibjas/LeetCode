func sortSentence(s string) string {
    arr := strings.Split(s, " ")
    str := make([]string, len(arr))

    for _, ele := range arr {
        num, _ := strconv.Atoi(string(ele[len(ele)-1]))
        str[num-1] = ele[:len(ele)-1]
    }
    return strings.Join(str," ")
}