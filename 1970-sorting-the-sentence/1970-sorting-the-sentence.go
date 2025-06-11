func sortSentence(s string) string {
    arr := strings.Split(s, " ")
    str := make([]string, len(arr))

    for _, ele := range arr {
        num := int(ele[len(ele)-1] - '1')
        str[num] = ele[:len(ele)-1]
    }
    return strings.Join(str," ")
}